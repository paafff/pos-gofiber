.PHONY: run seed reset-db

# Jalankan aplikasi
run:
	go run cmd/web/main.go

# Seed data dummy
seed:
	go run cmd/web/main.go seed

# Reset database
reset-db:
	go run cmd/web/main.go reset-db


# clean-run (reset database, seed data, and run application)
clean-run: reset-db seed run	