---
title: HR API
language_tabs:
  - go: Go

---

# HR API

Base URLs:

* <a href="http://localhost:8081">Testing Env: http://localhost:8081</a>

# Authentication

# HR API/Employee

## POST Create Employee

POST /api/v1/employee

> Body Parameters

```json
{
  "employee_name": "test_user_3",
  "employee_birth": "1993-11-12",
  "employee_role": "rd",
  "employee_email": "rd3@gmail.com"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» employee_name|body|string| yes |employee name|
|» employee_birth|body|string| yes |employee birth|
|» employee_role|body|string| yes |員工職位|
|» employee_email|body|string| yes |信箱|

> Response Examples

```json
{
  "data": {
    "employee_id": 3,
    "employee_name": "test_user_3",
    "employee_birth": "1993-11-12",
    "employee_role": "rd",
    "employee_email": "rd3@gmail.com"
  },
  "status": {
    "code": 200,
    "message": "OK"
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|»» employee_id|integer|true|none||none|
|»» employee_name|string|true|none||none|
|»» employee_birth|string|true|none||none|
|»» employee_role|string|true|none||none|
|»» employee_email|string|true|none||none|
|» status|object|true|none||none|
|»» code|integer|true|none||none|
|»» message|string|true|none||none|

## GET Query Employee

GET /api/v1/employee/{id}

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|id|path|string| yes |員工ID|

> Response Examples

```json
{
  "data": {
    "employee_id": 1,
    "employee_name": "test_user_1",
    "employee_birth": "1993-11-12",
    "employee_role": "rd",
    "employee_email": "rd@gmail.com"
  },
  "status": {
    "code": 200,
    "message": "OK"
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|»» employee_id|integer|true|none||none|
|»» employee_name|string|true|none||none|
|»» employee_birth|string|true|none||none|
|»» employee_role|string|true|none||none|
|»» employee_email|string|true|none||none|
|» status|object|true|none||none|
|»» code|integer|true|none||none|
|»» message|string|true|none||none|

## PUT Update Employee

PUT /api/v1/employee/{id}

> Body Parameters

```json
{
  "employee_name": "string",
  "employee_birth": "string",
  "employee_role": "string",
  "employee_email": "string"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|id|path|string| yes |none|
|body|body|object| no |none|
|» employee_name|body|string¦null| no |none|
|» employee_birth|body|string¦null| no |none|
|» employee_role|body|string¦null| no |none|
|» employee_email|body|string¦null| no |none|

> Response Examples

```json
{
  "data": {
    "employee_id": 1,
    "employee_name": "Rodney Jacobi",
    "employee_birth": "1993-11-12",
    "employee_role": "rd",
    "employee_email": "Trycia_Grady@gmail.com"
  },
  "status": {
    "code": 200,
    "message": "OK"
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|»» employee_id|integer|true|none||none|
|»» employee_name|string|true|none||none|
|»» employee_birth|string|true|none||none|
|»» employee_role|string|true|none||none|
|»» employee_email|string|true|none||none|
|» status|object|true|none||none|
|»» code|integer|true|none||none|
|»» message|string|true|none||none|

## DELETE Delete Employee

DELETE /api/v1/employee/{id}

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|id|path|string| yes |none|

> Response Examples

```json
{
  "data": {
    "employee_id": 1,
    "status": "deleted"
  },
  "status": {
    "code": 200,
    "message": "OK"
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|»» employee_id|integer|true|none||none|
|»» status|string|true|none||none|
|» status|object|true|none||none|
|»» code|integer|true|none||none|
|»» message|string|true|none||none|

## GET List Employees

GET /api/v1/employees

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|page|query|integer| no |頁數|
|limit|query|integer| no |單頁個數|
|field|query|string| no |排序欄位|
|order|query|string| no |asc or desc|

> Response Examples

```json
{
  "data": {
    "employee_list": [
      {
        "employee_id": 4,
        "employee_name": "test_user_4",
        "employee_birth": "1993-11-12",
        "employee_role": "rd",
        "employee_email": "rd4@gmail.com"
      },
      {
        "employee_id": 3,
        "employee_name": "test_user_3",
        "employee_birth": "1993-11-12",
        "employee_role": "rd",
        "employee_email": "rd3@gmail.com"
      },
      {
        "employee_id": 2,
        "employee_name": "test_user_2",
        "employee_birth": "1993-11-12",
        "employee_role": "rd",
        "employee_email": "rd2@gmail.com"
      }
    ],
    "pagination": {
      "page": 0,
      "limit": 0,
      "total": 0,
      "total_page": 0
    }
  },
  "status": {
    "code": 200,
    "message": "OK"
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|»» employee_list|[object]|true|none||none|
|»»» employee_id|integer|true|none||none|
|»»» employee_name|string|true|none||none|
|»»» employee_birth|string|true|none||none|
|»»» employee_role|string|true|none||none|
|»»» employee_email|string|true|none||none|
|»» pagination|object|true|none||none|
|»»» page|integer|true|none||none|
|»»» limit|integer|true|none||none|
|»»» total|integer|true|none||none|
|»»» total_page|integer|true|none||none|
|» status|object|true|none||none|
|»» code|integer|true|none||none|
|»» message|string|true|none||none|

# HR API/Department

## POST Create Department

POST /api/v1/department

> Body Parameters

```json
{
  "department_name": "dp1"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» department_name|body|string| yes |none|

> Response Examples

```json
{
  "data": {
    "department_id": 3,
    "department_name": "dp1"
  },
  "status": {
    "code": 200,
    "message": "OK"
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|»» department_id|integer|true|none||none|
|»» department_name|string|true|none||none|
|» status|object|true|none||none|
|»» code|integer|true|none||none|
|»» message|string|true|none||none|

## GET Query Department

GET /api/v1/department/{id}

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|id|path|string| yes |none|

> Response Examples

```json
{
  "data": {
    "department_id": 1,
    "department_name": "dp1"
  },
  "status": {
    "code": 200,
    "message": "OK"
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|»» department_id|integer|true|none||none|
|»» department_name|string|true|none||none|
|» status|object|true|none||none|
|»» code|integer|true|none||none|
|»» message|string|true|none||none|

## DELETE Delete Department

DELETE /api/v1/department/{id}

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|id|path|string| yes |none|

> Response Examples

```json
{
  "data": {
    "department_id": 1,
    "status": "deleted"
  },
  "status": {
    "code": 200,
    "message": "OK"
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|»» department_id|integer|true|none||none|
|»» status|string|true|none||none|
|» status|object|true|none||none|
|»» code|integer|true|none||none|
|»» message|string|true|none||none|

## PUT Update Department

PUT /api/v1/department/{id}

> Body Parameters

```json
{
  "department_name": "string"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|id|path|string| yes |none|
|body|body|object| no |none|
|» department_name|body|string¦null| no |none|

> Response Examples

```json
{
  "data": {
    "department_id": 1,
    "department_name": "Beulah Gutkowski"
  },
  "status": {
    "code": 200,
    "message": "OK"
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|»» department_id|integer|true|none||none|
|»» department_name|string|true|none||none|
|» status|object|true|none||none|
|»» code|integer|true|none||none|
|»» message|string|true|none||none|

## GET List Departments

GET /api/v1/departments

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|page|query|string| no |頁數|
|limit|query|integer| no |單頁個數|
|field|query|string| no |排序欄位|
|order|query|string| no |asc or desc|

> Response Examples

```json
{
  "data": {
    "department_list": [
      {
        "department_id": 2,
        "department_name": "dp1"
      },
      {
        "department_id": 3,
        "department_name": "dp1"
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 0,
      "total_page": 0
    }
  },
  "status": {
    "code": 200,
    "message": "OK"
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|»» department_list|[object]|true|none||none|
|»»» department_id|integer|true|none||none|
|»»» department_name|string|true|none||none|
|»» pagination|object|true|none||none|
|»»» page|integer|true|none||none|
|»»» limit|integer|true|none||none|
|»»» total|integer|true|none||none|
|»»» total_page|integer|true|none||none|
|» status|object|true|none||none|
|»» code|integer|true|none||none|
|»» message|string|true|none||none|

# HR API/Attendance

## POST Create Attendance

POST /api/v1/attendance

> Body Parameters

```json
{
  "employee_id": 0,
  "start_time": " 2006-01-02T15:04:05Z",
  "end_time": "2006-01-02T15:04:05Z",
  "type": "in"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» employee_id|body|integer| yes |員工ID|
|» start_time|body|string(date-time)| yes |RFC3339格式 ex: 2006-01-02T15:04:05Z|
|» end_time|body|string(date-time)| yes |RFC3339格式 ex: 2006-01-02T15:04:05Z|
|» type|body|string| yes |in or out|

#### Enum

|Name|Value|
|---|---|
|» type|in|
|» type|out|

> Response Examples

```json
{
  "data": {
    "attendance_id": 8,
    "employee_info": {
      "employee_id": 1,
      "employee_name": "test_user_1",
      "employee_birth": "1993-11-12",
      "employee_role": "rd",
      "employee_email": "rd3@gmail.com"
    },
    "attendance_in": "2025-02-01T00:00:00Z",
    "attendance_out": "2025-02-01T10:00:00Z"
  },
  "status": {
    "code": 200,
    "message": "OK"
  }
}
```

```json
{
  "data": {
    "attendance_id": 7,
    "employee_info": {
      "employee_id": 1,
      "employee_name": "test_user_1",
      "employee_birth": "1993-11-12",
      "employee_role": "rd",
      "employee_email": "rd3@gmail.com"
    },
    "attendance_in": "2025-02-02T00:00:00Z",
    "attendance_out": "2025-02-02T10:00:00Z"
  },
  "status": {
    "code": 200,
    "message": "OK"
  }
}
```

```json
{
  "data": {
    "attendance_id": 8,
    "employee_info": {
      "employee_id": 1,
      "employee_name": "test_user_1",
      "employee_birth": "1993-11-12",
      "employee_role": "rd",
      "employee_email": "rd3@gmail.com"
    },
    "attendance_in": "2025-02-01T00:00:00Z",
    "attendance_out": "2025-02-01T10:00:00Z"
  },
  "status": {
    "code": 200,
    "message": "OK"
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|»» attendance_id|integer|true|none||none|
|»» employee_info|object|true|none||none|
|»»» employee_id|integer|true|none||none|
|»»» employee_name|string|true|none||none|
|»»» employee_birth|string|true|none||none|
|»»» employee_role|string|true|none||none|
|»»» employee_email|string|true|none||none|
|»» attendance_in|string|true|none||none|
|»» attendance_out|string|true|none||none|
|» status|object|true|none||none|
|»» code|integer|true|none||none|
|»» message|string|true|none||none|

## GET Query Attendance

GET /api/v1/attendance/{id}

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|id|path|string| yes |簽到ID|

> Response Examples

```json
{
  "data": {
    "attendance_id": 5,
    "employee_info": {
      "employee_id": 1,
      "employee_name": "test_user_1",
      "employee_birth": "1993-11-12",
      "employee_role": "rd",
      "employee_email": "rd3@gmail.com"
    },
    "attendance_in": "2025-02-03T00:00:00Z",
    "attendance_out": "2025-02-03T10:00:00Z"
  },
  "status": {
    "code": 200,
    "message": "OK"
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|»» attendance_id|integer|true|none||none|
|»» employee_info|object|true|none||none|
|»»» employee_id|integer|true|none||none|
|»»» employee_name|string|true|none||none|
|»»» employee_birth|string|true|none||none|
|»»» employee_role|string|true|none||none|
|»»» employee_email|string|true|none||none|
|»» attendance_in|string|true|none||none|
|»» attendance_out|string|true|none||none|
|» status|object|true|none||none|
|»» code|integer|true|none||none|
|»» message|string|true|none||none|

## PUT Update Attendance

PUT /api/v1/attendance/{id}

> Body Parameters

```json
{
  "start_time": "2019-08-24T14:15:22Z",
  "end_time": "2019-08-24T14:15:22Z",
  "type": "string"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|id|path|string| yes |none|
|body|body|object| no |none|
|» start_time|body|string(date-time)¦null| no |RFC3339|
|» end_time|body|string(date-time)¦null| no |RFC3339|
|» type|body|string¦null| no |in or out|

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

## DELETE Delete Attendance

DELETE /api/v1/attendance/{id}

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|id|path|string| yes |none|

> Response Examples

```json
{
  "data": {
    "attendance_id": 5,
    "Status": "deleted"
  },
  "status": {
    "code": 200,
    "message": "OK"
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|»» attendance_id|integer|true|none||none|
|»» Status|string|true|none||none|
|» status|object|true|none||none|
|»» code|integer|true|none||none|
|»» message|string|true|none||none|

## GET List Attendances

GET /api/v1/attendances

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|page|query|integer| no |頁數|
|limit|query|integer| no |單頁個數|
|field|query|string| no |排序欄位|
|order|query|string| no |asc or desc|

> Response Examples

```json
{
  "data": {
    "attendance_list": [
      {
        "attendance_id": 6,
        "employee_info": {
          "employee_id": 1,
          "employee_name": "test_user_1",
          "employee_birth": "1993-11-12",
          "employee_role": "rd",
          "employee_email": "rd3@gmail.com"
        },
        "attendance_in": "2025-02-04T00:00:00Z",
        "attendance_out": "2025-02-04T10:00:00Z"
      },
      {
        "attendance_id": 7,
        "employee_info": {
          "employee_id": 1,
          "employee_name": "test_user_1",
          "employee_birth": "1993-11-12",
          "employee_role": "rd",
          "employee_email": "rd3@gmail.com"
        },
        "attendance_in": "2025-02-02T00:00:00Z",
        "attendance_out": "2025-02-02T10:00:00Z"
      },
      {
        "attendance_id": 8,
        "employee_info": {
          "employee_id": 1,
          "employee_name": "test_user_1",
          "employee_birth": "1993-11-12",
          "employee_role": "rd",
          "employee_email": "rd3@gmail.com"
        },
        "attendance_in": "2025-02-01T00:00:00Z",
        "attendance_out": "2025-02-01T10:00:00Z"
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 0,
      "total_page": 0
    }
  },
  "status": {
    "code": 200,
    "message": "OK"
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|»» attendance_list|[object]|true|none||none|
|»»» attendance_id|integer|true|none||none|
|»»» employee_info|object|true|none||none|
|»»»» employee_id|integer|true|none||none|
|»»»» employee_name|string|true|none||none|
|»»»» employee_birth|string|true|none||none|
|»»»» employee_role|string|true|none||none|
|»»»» employee_email|string|true|none||none|
|»»» attendance_in|string|true|none||none|
|»»» attendance_out|string|true|none||none|
|»» pagination|object|true|none||none|
|»»» page|integer|true|none||none|
|»»» limit|integer|true|none||none|
|»»» total|integer|true|none||none|
|»»» total_page|integer|true|none||none|
|» status|object|true|none||none|
|»» code|integer|true|none||none|
|»» message|string|true|none||none|

# HR API/Leave

## POST Create Leave

POST /api/v1/leave

> Body Parameters

```json
{
  "employee_id": 0,
  "start_time": "string",
  "end_time": "string",
  "type": "annual"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» employee_id|body|integer| yes |員工ID|
|» start_time|body|string| yes |開始時間|
|» end_time|body|string| yes |結束時間|
|» type|body|string| yes |類型|

#### Enum

|Name|Value|
|---|---|
|» type|annual|
|» type|sick|
|» type|unpaid|

> Response Examples

```json
{
  "data": {
    "leave_id": 3,
    "employee_info": {
      "employee_id": 2,
      "employee_name": "test_user_2",
      "employee_birth": "1993-11-12",
      "employee_role": "rd",
      "employee_email": "rd3@gmail.com"
    },
    "leave_start": "2025-02-03T20:49:08.169Z",
    "leave_end": "2025-02-03T11:11:27.29Z"
  },
  "status": {
    "code": 200,
    "message": "OK"
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|»» leave_id|integer|true|none||none|
|»» employee_info|object|true|none||none|
|»»» employee_id|integer|true|none||none|
|»»» employee_name|string|true|none||none|
|»»» employee_birth|string|true|none||none|
|»»» employee_role|string|true|none||none|
|»»» employee_email|string|true|none||none|
|»» leave_start|string|true|none||none|
|»» leave_end|string|true|none||none|
|» status|object|true|none||none|
|»» code|integer|true|none||none|
|»» message|string|true|none||none|

## GET Query Leave

GET /api/v1/leave/{id}

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|id|path|string| yes |none|

> Response Examples

```json
{
  "data": {
    "leave_id": 1,
    "employee_info": {
      "employee_id": 2,
      "employee_name": "test_user_2",
      "employee_birth": "1993-11-12",
      "employee_role": "rd",
      "employee_email": "rd3@gmail.com"
    },
    "leave_start": "2025-02-03T20:49:08.169Z",
    "leave_end": "2025-02-03T11:11:27.29Z"
  },
  "status": {
    "code": 200,
    "message": "OK"
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|»» leave_id|integer|true|none||none|
|»» employee_info|object|true|none||none|
|»»» employee_id|integer|true|none||none|
|»»» employee_name|string|true|none||none|
|»»» employee_birth|string|true|none||none|
|»»» employee_role|string|true|none||none|
|»»» employee_email|string|true|none||none|
|»» leave_start|string|true|none||none|
|»» leave_end|string|true|none||none|
|» status|object|true|none||none|
|»» code|integer|true|none||none|
|»» message|string|true|none||none|

## PUT Update Leave

PUT /api/v1/leave/{id}

> Body Parameters

```json
{
  "start_time": "string",
  "end_time": "string",
  "type": "string"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|id|path|string| yes |none|
|body|body|object| no |none|
|» start_time|body|string¦null| no |none|
|» end_time|body|string¦null| no |none|
|» type|body|string¦null| no |none|

> Response Examples

```json
{
  "data": {
    "leave_id": 1,
    "employee_info": {
      "employee_id": 2,
      "employee_name": "test_user_2",
      "employee_birth": "1993-11-12",
      "employee_role": "rd",
      "employee_email": "rd3@gmail.com"
    },
    "leave_start": "2025-02-04T04:31:43.828Z",
    "leave_end": "2025-02-03T15:46:04.69Z"
  },
  "status": {
    "code": 200,
    "message": "OK"
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|»» leave_id|integer|true|none||none|
|»» employee_info|object|true|none||none|
|»»» employee_id|integer|true|none||none|
|»»» employee_name|string|true|none||none|
|»»» employee_birth|string|true|none||none|
|»»» employee_role|string|true|none||none|
|»»» employee_email|string|true|none||none|
|»» leave_start|string|true|none||none|
|»» leave_end|string|true|none||none|
|» status|object|true|none||none|
|»» code|integer|true|none||none|
|»» message|string|true|none||none|

## DELETE Delete Leave

DELETE /api/v1/leave/{id}

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|id|path|string| yes |none|

> Response Examples

```json
{
  "data": {
    "leave_id": 1,
    "status": "deleted"
  },
  "status": {
    "code": 200,
    "message": "OK"
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|»» leave_id|integer|true|none||none|
|»» status|string|true|none||none|
|» status|object|true|none||none|
|»» code|integer|true|none||none|
|»» message|string|true|none||none|

## GET List Leaves

GET /api/v1/leaves

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|page|query|integer| no |頁數|
|limit|query|integer| no |單頁個數|
|field|query|string| no |排序欄位|
|order|query|string| no |asc or desc|

> Response Examples

```json
{
  "data": {
    "leave_list": [
      {
        "leave_id": 2,
        "employee_info": {
          "employee_id": 2,
          "employee_name": "test_user_2",
          "employee_birth": "1993-11-12",
          "employee_role": "rd",
          "employee_email": "rd3@gmail.com"
        },
        "leave_start": "2025-02-03T20:49:08.169Z",
        "leave_end": "2025-02-03T11:11:27.29Z"
      },
      {
        "leave_id": 3,
        "employee_info": {
          "employee_id": 2,
          "employee_name": "test_user_2",
          "employee_birth": "1993-11-12",
          "employee_role": "rd",
          "employee_email": "rd3@gmail.com"
        },
        "leave_start": "2025-02-03T20:49:08.169Z",
        "leave_end": "2025-02-03T11:11:27.29Z"
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 0,
      "total_page": 0
    }
  },
  "status": {
    "code": 200,
    "message": "OK"
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

# HR API/Statistics

## POST Statistics

POST /api/v1/attendance-statistic

> Body Parameters

```json
{
  "employee_id_list": [
    0
  ],
  "start": "2019-08-24T14:15:22Z",
  "end": "2019-08-24T14:15:22Z"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» employee_id_list|body|[integer]| yes |none|
|» start|body|string(date-time)| yes |none|
|» end|body|string(date-time)| yes |none|

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

# Data Schema

<h2 id="tocS_Paginator">Paginator</h2>

<a id="schemapaginator"></a>
<a id="schema_Paginator"></a>
<a id="tocSpaginator"></a>
<a id="tocspaginator"></a>

```json
{
  "page": 0,
  "limit": 0
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|page|integer|true|none||第幾頁|
|limit|integer|true|none||每頁個數|

<h2 id="tocS_Pet">Pet</h2>

<a id="schemapet"></a>
<a id="schema_Pet"></a>
<a id="tocSpet"></a>
<a id="tocspet"></a>

```json
{
  "id": 1,
  "category": {
    "id": 1,
    "name": "string"
  },
  "name": "doggie",
  "photoUrls": [
    "string"
  ],
  "tags": [
    {
      "id": 1,
      "name": "string"
    }
  ],
  "status": "available"
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|id|integer(int64)|true|none||Pet ID|
|category|[Category](#schemacategory)|true|none||group|
|name|string|true|none||name|
|photoUrls|[string]|true|none||image URL|
|tags|[[Tag](#schematag)]|true|none||tag|
|status|string|true|none||Pet Sales Status|

#### Enum

|Name|Value|
|---|---|
|status|available|
|status|pending|
|status|sold|

<h2 id="tocS_Category">Category</h2>

<a id="schemacategory"></a>
<a id="schema_Category"></a>
<a id="tocScategory"></a>
<a id="tocscategory"></a>

```json
{
  "id": 1,
  "name": "string"
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|id|integer(int64)|false|none||Category ID|
|name|string|false|none||Category Name|

<h2 id="tocS_Tag">Tag</h2>

<a id="schematag"></a>
<a id="schema_Tag"></a>
<a id="tocStag"></a>
<a id="tocstag"></a>

```json
{
  "id": 1,
  "name": "string"
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|id|integer(int64)|false|none||Tag ID|
|name|string|false|none||Tag Name|

