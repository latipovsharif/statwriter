FROM reg.techhprof.ru/profitclicks/base-ubuntu:latest
WORKDIR /app
COPY ./dailystatuploader ./
COPY ./config.yaml ./
ENV ENV=prod
CMD ["./dailystatuploader"]