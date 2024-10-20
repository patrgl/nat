
# nat

Fast & simple concurrent api-based file storage with token authentication made in Go.

Multiple actions within one request, say you are attempting to upload or download multiple files, will be done concurrently in order to speed up operations.

## API Reference

### Authorization token example

{Authorization : Bearer `Your token`}

### Upload a file

```http
  PUT /api/upload
```
Header:
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `Authorization` | `string` | **Required**.  |

Body:

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `file`      | `file` | **Required**. File to be uploaded |

### Upload multiple files

```http
  PUT /api/multi-upload
```
Header:
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `Authorization` | `string` | **Required**.  |

Body:

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `file`      | `file` | **Required**. First file |
| `file`      | `file` | **Required**. Second file |
| `file`      | `file` | **Required**. ... file |

### Download one file

```http
  GET /api/download/{filename}
```

Header:
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `Authorization` | `string` | **Required**.  |


### Delete one file

```http
  DELETE /api/delete/{filename}
```

Header:
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `Authorization` | `string` | **Required**.  |

### Download multiple files
Downloads a zip of the requested files.

```http
  GET /api/multi-download?file=file1.jpg&file=file2.png...
```

Header:
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `Authorization` | `string` | **Required**.  |

Note: Files not found will be ingored, and their names will be added to the response header notifiying you they were not found.

### Delete multiple files

```http
  DELETE /api/multi-delete?file=file1.jpg&file=file2.png...
```

Header:
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `Authorization` | `string` | **Required**.  |

Note: Files not found will be ingored, and their names will be added to the response header notifiying you they were not found.

## Deployment

To deploy this project run

```bash
  go run main.go
```

Remeber to set a the "token" env variable to be a proper randomized token before using this outside of testing, and removing os.Setenv("thisisatoken") in main.go, once you do. "thisisatoken" is the default token set by this line.
