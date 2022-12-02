FROM golang:1.19rc1-stretch

# sqlboilerのインストール
RUN go install github.com/volatiletech/sqlboiler/v4@v4.8.6
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@v4.8.6

# 設定ファイルのコピー
COPY ./sqlboiler.local.toml /tmp/conf/sqlboiler.toml

COPY ./sqlboiler.sh /tmp/sqlboiler.sh
RUN chmod a+x /tmp/sqlboiler.sh

RUN mkdir /tmp/modules
RUN mkdir /tmp/modules/dbmodels

# run --rm での実行を想定
ENTRYPOINT ["/bin/sh", "-c"]