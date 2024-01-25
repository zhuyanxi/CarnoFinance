build-front:
	cd frontend && npm run build && cp -r build/* ../public