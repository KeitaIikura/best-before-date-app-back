DIR=${SQLBOILER_OUT_DIR}

# ファイルの生成
sqlboiler mysql -o /tmp/dbmodels -c /tmp/conf/sqlboiler.toml

# マウントパスへの移動
cp -rp /tmp/dbmodels/* ${DIR}
