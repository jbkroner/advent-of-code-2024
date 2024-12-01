.PHONY: install install-dev lint format pre-commit

install:
	pip install .

install-dev:
	pip install .[dev]

lint:
	ruff check .

format:
	black .
	ruff check --fix .

pre-commit:
	pre-commit run --all-files

docker-build:
	docker build -t turbo:dev_latest .

act-x86:
	act -j pre-commit -j build

act-arm:
	act -j pre-commit -j build --container-architecture linux/arm64

setup: install-dev
	pre-commit install

clean:
	find . -type d -name "__pycache__" -exec rm -rf {} +
	find . -type f -name "*.pyc" -delete