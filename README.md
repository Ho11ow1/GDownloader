# GDownloader
[![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-blue.svg)](#)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/license/MIT)<br/>
**An easy to use and effective online archive downloader.**<br/>

---
## Features

- **Flexible Input**
   - Single URL mode — pass a direct link via `-url`
   - Batch mode — provide a plain text file of URLs via `-file`, one per line

- **Download Filtering**
   - Limit the number of downloaded files with `-limit`
   - Filter by filename prefix with `-prefix`
   - Filter by file extension with `-extension`

- **Concurrent Downloads**
   - Each supported URL is dispatched to its handler in a separate goroutine, all synchronized with a `WaitGroup`

- **Service Architecture**
   - Pluggable `IDownloadProvider` interface — new services are registered in `AvailableServices` and are auto-matched to URLs via `Supports()`
   
- **Supported Services** — single file and full album downloads across all providers
   - [Bunkr](https://bunkr.site)
   - [Filester](https://filester.gg)
   - [FileDitch](https://fileditch.com)

---
## Usage

```
go run . [flags]
```

| Flag | Type | Description |
|---|---|---|
| `-url` | string | Single URL to download from |
| `-file` | string | Path to a file containing one URL per line |
| `-limit` | uint | Stop after downloading this many files |
| `-prefix` | string | Only download files whose name starts with this string |
| `-extension` | string | Only download files with this extension |

> `-url` and `-file` are mutually exclusive — provide one, not both.

> **Note:** `-limit`, `-prefix`, and `-extension` are not yet implemented.

**Single URL**
```
GDownloader -url "https://bunkr.site/a/example"
```

**Batch file**
```
GDownloader -file "urls.txt"
```

**With filters**
```
GDownloader -url "https://bunkr.site/a/example" -limit 10 -prefix "chapter_" -extension ".jpg"
```

**urls.txt format**
```
https://bunkr.site/a/example
https://fileditch.com/file/example
https://filester.gg/file/example
```


---
## Requirements
- [GoLang](https://go.dev/dl/) binary version 1.26.3 or higher
- Any kind of command prompt (cmd, ps, etc...)

---
## License

MIT License - see [LICENSE](LICENSE)

If you find any issues during usage, please create a github issue [Here](https://github.com/Ho11ow1/GDownloader/issues)
