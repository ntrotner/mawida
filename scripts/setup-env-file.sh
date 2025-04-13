target_file="./environment/global.env"
source_files=("backend.env" "database.env" "nginx.env" "ui.env")
truncate -s 0 "$target_file"

cat "./environment/docker.env" >> "$target_file"
echo "" >> "$target_file"

for source_file in "${source_files[@]}"; do
  echo "./environment/$1/$source_file"

  if [ -f "./environment/$1/$source_file" ]; then
    cat "./environment/$1/$source_file" >> "$target_file"
    echo "" >> "$target_file"
  fi
done

