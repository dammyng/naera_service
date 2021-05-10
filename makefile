docker_rebuild:
	cd naera && docker compose down && docker compose up --build -V -d