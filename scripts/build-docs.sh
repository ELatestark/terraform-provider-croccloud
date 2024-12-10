#!/bin/bash

# Генерация документации
cd website
mkdocs build --clean

# Путь к директории, из которой необходимо загружать файлы
SOURCE_DIR="website/site/"

# S3 бакет и путь
S3_BUCKET="s3://doc-for-terraform/"

# Функция для загрузки всех остальных файлов без изменения Content-Type
upload_other_files () {
    s3cmd sync "$SOURCE_DIR" "$S3_BUCKET" --acl-public
}

# Функция для загрузки файлов с заданным Content-Type
upload_css_files () {
    find "$SOURCE_DIR" -type f -name "*.css" | while read -r file; do
    s3cmd modify "$S3_BUCKET${file#$SOURCE_DIR}" --add-header='Content-Type:text/css'
done
}

# Выполнение функций
cd ..
upload_other_files
upload_css_files
