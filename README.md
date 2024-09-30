
# nat

Fast & simple api-based file storage with token authentication made in Go.

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
