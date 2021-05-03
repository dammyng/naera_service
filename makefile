docker_rebuild:
	cd docker && docker compose down && docker compose up --build -V -d