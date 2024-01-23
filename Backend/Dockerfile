FROM python:3.11-slim

WORKDIR /app

COPY requirements.txt requirements.txt

RUN pip install --upgrade pip

RUN pip install --default-timeout=1000 -r requirements.txt

RUN apt-get update && apt-get install -y libgl1-mesa-glx libglib2.0-0

COPY . .

CMD ["python", "main.py"]

EXPOSE 80

FROM python:3.11-slim

WORKDIR /app

COPY requirements.txt requirements.txt

RUN pip install --upgrade pip

RUN apt-get update && \
    apt-get install -y libgl1-mesa-glx libglib2.0-0 \
    gcc python3-dev

RUN pip install --default-timeout=1000 -r requirements.txt

COPY . .

CMD ["python", "main.py"]

EXPOSE 80
