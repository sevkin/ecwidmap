# Creates sitemap.xml for Ecwid shop

TODO: english readme (you can help me to translate it)

Запускается из командной строки, при запуске требуется указать:
- storeID - ID магазина Ecwid
- token - токен с правами доступа на чтение товаров и категорий

- daily - необязательный параметр, добавляет `<changefreq>daily</changefreq>` ко всем ссылкам

`storeID` и `token` можно указать как через переменные окружения, так и через параметры коммандной строки. Параметры имеют приоритет перед окружением. Пример:

```shell
export ECWID_STOREID=123456
export ECWID_TOKEN=secret_token

ecwidmap >/path/to/website/root/sitemap.xml
```

```shell
ecwidmap -storeid 123456 -token secret_token >/path/to/website/root/sitemap.xml
```

`sitemap.xml` выводится на `/dev/stdout` который нужно **перенаправить** в файл.