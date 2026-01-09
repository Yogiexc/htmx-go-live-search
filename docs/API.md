# API Documentation

## Endpoints

### `GET /`
**Description:** Render main search page

**Response:** HTML page with search input

**Example:**
```
http://localhost:8080/
```

---

### `GET /search`
**Description:** Search users by query parameter

**Query Parameters:**
- `q` (string, required): Search query

**Response:** HTML fragment with search results

**Example:**
```
http://localhost:8080/search?q=jakarta
```

**Response Format:**
```html
<div class="results-header">
    <p>Ditemukan <strong>2</strong> hasil...</p>
</div>
<div class="results-list">
    <!-- User cards -->
</div>
```

---

## Data Structure

### User
```go
type User struct {
    ID    int
    Name  string
    Email string
    City  string
}
```

### Search Logic
- Case-insensitive search
- Matches in: Name, Email, City
- Returns partial matches
- In-memory filtering