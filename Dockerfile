FROM reg.techhprof.ru/profitclicks/base-ubuntu:latest
WORKDIR /app
COPY ./dailystatuploader ./
COPY ./config.yaml ./
COPY ./date_stat.xlsx ./
COPY ./country_stat.xlsx ./
COPY ./device_pc_stat.xlsx ./
COPY ./device_mobile_stat.xlsx ./
COPY ./subscription_stat.xlsx ./
ENV ENV=prod
CMD ["./dailystatuploader"]