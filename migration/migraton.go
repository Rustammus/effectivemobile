package migration

import (
	"EffectiveMobile/internal/config"
	"EffectiveMobile/pkg/client/postgres"
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type migration struct {
	QueryStr  string
	Name      string
	Id        int
	isApplied bool
}

func DoMigrate() {
	conf := config.GetConfig()
	conn, err := postgres.NewClient(context.TODO(), conf.Storage)
	if err != nil {
		log.Fatal(err)
		return
	}

	createMigrationTable(conn)

	files, err := os.ReadDir("C:/dev/projects/EffectiveMobile/migration/migrations")
	if err != nil {
		log.Fatalf("Failed to read migrations directory: %v\n", err)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	migratons := getMigrations(files)
	doMigrations(migratons, conn)
}

func createMigrationTable(conn *pgx.Conn) {
	q := `CREATE TABLE IF NOT EXISTS public.migrations (
			id INT PRIMARY KEY,
			name TEXT NOT NULL,
			applied_at timestamptz DEFAULT CURRENT_TIMESTAMP
		)`
	_, err := conn.Exec(context.TODO(), q)
	if err != nil {
		log.Fatalf("Failed to create migrations table: %v\n", err)
	}

}

func getMigrations(files []os.DirEntry) []migration {
	migratons := make([]migration, 0)
	ids := make(map[int]bool)
	for i := 0; i < len(files); i++ {
		substrs := strings.Split(files[i].Name(), "_")
		if len(substrs) < 3 {
			log.Fatalf("Uncorrect migration name: %s. Example: '0000_create_some_up.sql'\n", files[i].Name())
		}

		idStr := substrs[0]
		isUp := false
		if substrs[len(substrs)-1] == "up.sql" {
			isUp = true
		} else if substrs[len(substrs)-1] != "down.sql" {
			log.Fatalf("migration file %s unknown state. Should be '_up.sql' or '_down.sql'\n", files[i].Name())
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatalf("migration file %s id parse error. %s \n", files[i].Name(), err)
		}
		if _, ok := ids[id]; ok {
			log.Fatalf("migration file %s id duplicate\n", files[i].Name())
		}

		readText, err := os.ReadFile("C:/dev/projects/EffectiveMobile/migration/migrations/" + files[i].Name())
		if err != nil {
			log.Fatalf("migration file %s read error. %s \n", files[i].Name(), err)
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

func doMigrations(m []migration, conn *pgx.Conn) {
	checkMigrations(m, conn)

	for i := 0; i < len(m); i++ {
		if m[i].isApplied {
			log.Printf("migration %s already applied, skipping\n", m[i].Name)
			continue
		}
		_, err := conn.Exec(context.TODO(), m[i].QueryStr)
		if err != nil {
			log.Fatalf("Failed to apply migration %s: %v\n", m[i].Name, err)
		}
		_, err = conn.Exec(context.TODO(), "INSERT INTO public.migrations (id, name) VALUES ($1, $2)", m[i].Id, m[i].Name)
		if err != nil {
			log.Fatalf("Failed to record migration %s: %v\n", m[i].Name, err)
		}
	}

}

func checkMigrations(m []migration, conn *pgx.Conn) {
	q := `SELECT id, name FROM public.migrations ORDER BY applied_at`
	rows, err := conn.Query(context.TODO(), q)
	if err != nil {
		log.Fatalf("Failed to read migrations: %v\n", err)
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		if i >= len(m) {
			log.Printf("WARN to many migration in db. Exepted %d but get more", len(m))
			break
		}

		readId := -1
		readName := ""

		err := rows.Scan(&readId, &readName)
		if err != nil {
			log.Fatalf("Failed to scan migration: %v\n", err)
			return
		}
		if readId == m[i].Id && readName != m[i].Name {
			log.Printf("WARN different migration name db: %s, file: %s", readName, m[i].Name)
		} else if readId != m[i].Id {
			log.Fatalf("migration id db mismatch in row %d. Expected %d, got %d.\n", i, m[i].Id, readId)
		}

		m[i].isApplied = true
	}
}
