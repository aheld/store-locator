dev:
	air -c ./.air.toml & \
  browser-sync start \
  --files '**/*.templ' \
  --port 8080 \
	--proxy 'http://127.0.0.1:8000' \
	--reload-delay 1500

godev:
	air -c ./.air.toml	

run:
	go run *.go

css:
	npx tailwindcss -i input.css -o assets/output.css --minify

htmx:
	curl "https://unpkg.com/htmx.org@1.9.10/dist/htmx.min.js" > assets/htmx_1.9.10.js

container-build:
	docker build . -t markets

container-run: container-build
	docker run -p8000:8000 markets  

deploy:
	flyctl deploy
