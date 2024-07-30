package migration

import (
	"EffectiveMobile/internal/config"
	"EffectiveMobile/pkg/client/postgres"
	"EffectiveMobile/pkg/logging"
	"context"
	"github.com/jackc/pgx/v5"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Migrator struct {
	logger logging.Logger
	conn   *pgx.Conn
}

type migration struct {
	QueryStr  string
	Name      string
	Id        int
	isApplied bool
}

func NewMigrator(logger logging.Logger) *Migrator {
	return &Migrator{logger: logger}
}

func (m *Migrator) Up() {
	conf := config.GetConfig()
	var err error
	m.conn, err = postgres.NewClient(context.TODO(), conf.Storage)
	if err != nil {
		m.logger.Fatal(err)
		return
	}
	defer m.conn.Close(context.Background())

	err = m.createMigrationTable()
	if err != nil {
		m.logger.Fatalf("Failed to create migrations table: %v\n", err)
	}

	files, err := os.ReadDir("C:/dev/projects/EffectiveMobile/migration/migrations")
	if err != nil {
		m.logger.Fatalf("Failed to read migrations directory: %v\n", err)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	migratons := m.getMigrations(files)
	m.doMigrations(migratons)
}

func (m *Migrator) createMigrationTable() error {
	q := `CREATE TABLE IF NOT EXISTS public.migrations (
			id INT PRIMARY KEY,
			name TEXT NOT NULL,
			applied_at timestamptz DEFAULT CURRENT_TIMESTAMP
		)`
	_, err := m.conn.Exec(context.TODO(), q)
	return err
}

func (m *Migrator) getMigrations(files []os.DirEntry) []migration {
	migratons := make([]migration, 0)
	ids := make(map[int]bool)
	for i := 0; i < len(files); i++ {
		substrs := strings.Split(files[i].Name(), "_")
		if len(substrs) < 3 {
			m.logger.Fatalf("Uncorrect migration name: %s. Example: '0000_create_some_up.sql'\n", files[i].Name())
		}

		idStr := substrs[0]
		isUp := false
		if substrs[len(substrs)-1] == "up.sql" {
			isUp = true
		} else if substrs[len(substrs)-1] != "down.sql" {
			m.logger.Fatalf("migration file %s unknown state. Must be '_up.sql' or '_down.sql'\n", files[i].Name())
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			m.logger.Fatalf("migration file %s id parse error. %s \n", files[i].Name(), err)
		}
		if _, ok := ids[id]; ok {
			m.logger.Fatalf("migration file %s id duplicate\n", files[i].Name())
		}

		readText, err := os.ReadFile("C:/dev/projects/EffectiveMobile/migration/migrations/" + files[i].Name())
		if err != nil {
			m.logger.Fatalf("migration file %s read error. %s \n", files[i].Name(), err)
		}

		if isUp {
			migratons = append(migratons, migration{
				QueryStr: string(readText),
				Name:     files[i].Name(),
				Id:       id,
			})
		}
	}
	return migratons
}

func (m *Migrator) doMigrations(migrations []migration) {
	m.checkMigrations(migrations)

	for i := 0; i < len(migrations); i++ {
		if migrations[i].isApplied {
			m.logger.Printf("migration %s already applied, skipping\n", migrations[i].Name)
			continue
		}
		_, err := m.conn.Exec(context.TODO(), migrations[i].QueryStr)
		if err != nil {
			m.logger.Fatalf("Failed to apply migration %s: %v\n", migrations[i].Name, err)
		}
		_, err = m.conn.Exec(context.TODO(), "INSERT INTO public.migrations (id, name) VALUES ($1, $2)", migrations[i].Id, migrations[i].Name)
		if err != nil {
			m.logger.Fatalf("Failed to record migration %s: %v\n", migrations[i].Name, err)
		}
	}

}

func (m *Migrator) checkMigrations(migrations []migration) {
	q := `SELECT id, name FROM public.migrations ORDER BY applied_at`
	rows, err := m.conn.Query(context.TODO(), q)
	if err != nil {
		m.logger.Fatalf("Failed to read migrations: %v\n", err)
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		if i >= len(migrations) {
			m.logger.Printf("WARN to many migration in db. Exepted %d but get more", len(migrations))
			break
		}

		readId := -1
		readName := ""

		err := rows.Scan(&readId, &readName)
		if err != nil {
			m.logger.Fatalf("Failed to scan migration: %v\n", err)
			return
		}
		if readId == migrations[i].Id && readName != migrations[i].Name {
			m.logger.Printf("WARN different migration name db: %s, file: %s", readName, migrations[i].Name)
		} else if readId != migrations[i].Id {
			m.logger.Fatalf("migration id db mismatch in row %d. Expected %d, got %d.\n", i, migrations[i].Id, readId)
		}

		migrations[i].isApplied = true
	}
}
