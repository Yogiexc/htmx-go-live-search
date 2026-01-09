# 🔍 HTMX Live Search with Go

> **Frontend doesn't always need to be a SPA.**  
> This project explores server-driven UI using HTMX and Go.

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![HTMX](https://img.shields.io/badge/HTMX-3D72D7?style=for-the-badge&logo=htmx&logoColor=white)
![HTML5](https://img.shields.io/badge/HTML5-E34F26?style=for-the-badge&logo=html5&logoColor=white)

## 🎯 Kenapa Project Ini?

Bosen sama overhead SPA? Capek setup Webpack, Babel, dan temen-temennya? Project ini bukti bahwa **kamu gak butuh React/Vue/Angular** untuk bikin interactive web app yang keren.

**HTMX** memungkinkan kita bikin **live search tanpa satu baris JavaScript pun** yang kita tulis sendiri. Semua logic di server, HTML dikirim langsung ke browser. Simple. Elegant. Powerful.

---

## ✨ Features

✅ **Live Search** - Real-time filtering tanpa reload page  
✅ **Zero JavaScript** - Cuma pakai HTMX (14KB library)  
✅ **Server-Rendered** - HTML di-render di backend  
✅ **Loading Indicator** - UX yang smooth dengan spinner  
✅ **Debouncing** - Request optimized dengan delay 300ms  
✅ **No Build Tools** - Gak perlu npm, webpack, babel, dll  

---

## 🏗️ Tech Stack

| Technology | Purpose |
|------------|---------|
| **Go (net/http)** | Backend server & routing |
| **html/template** | Server-side HTML rendering |
| **HTMX** | Frontend interactivity |
| **Vanilla CSS** | Styling (no framework) |

---

## 📂 Project Structure

```
htmx-go-live-search/
├── main.go              # Backend logic & routes
├── go.mod               # Go module definition
├── templates/
│   ├── index.html       # Main page
│   └── results.html     # Search results fragment
└── static/
    └── style.css        # Styling
```

---

## 🚀 Quick Start

### Prerequisites
- Go 1.21+ installed
- A web browser

### Installation & Run

```bash
# Clone repository
git clone https://github.com/YOUR_USERNAME/htmx-go-live-search.git
cd htmx-go-live-search

# Initialize Go module
go mod init htmx-go-live-search

# Run server
go run main.go
```

Server akan jalan di **http://localhost:8080** 🎉

---

## 🎨 How It Works

### 1️⃣ User Types in Search Box

```html
<input 
    type="text" 
    name="q"
    hx-get="/search"
    hx-trigger="keyup changed delay:300ms"
    hx-target="#search-results"
    hx-indicator="#loading"
>
```

### 2️⃣ HTMX Sends Request

HTMX otomatis:
- Detect event `keyup`
- Wait 300ms (debouncing)
- Send GET request ke `/search?q=jakarta`
- Show loading indicator

### 3️⃣ Server Filters Data

```go
func handleSearch(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query().Get("q")
    
    // Filter users in-memory
    for _, user := range users {
        if strings.Contains(user.Name, query) {
            results = append(results, user)
        }
    }
    
    // Render HTML fragment
    templates.ExecuteTemplate(w, "results.html", results)
}
```

### 4️⃣ HTMX Updates DOM

HTMX inject HTML fragment ke `#search-results` div. **No page reload. No JavaScript needed.**

---

## 🔥 HTMX Attributes Explained

| Attribute | Purpose | Example |
|-----------|---------|---------|
| `hx-get` | HTTP GET request | `hx-get="/search"` |
| `hx-trigger` | Event yang trigger request | `hx-trigger="keyup changed delay:300ms"` |
| `hx-target` | Element yang di-update | `hx-target="#search-results"` |
| `hx-indicator` | Loading indicator | `hx-indicator="#loading"` |
| `hx-swap` | Cara replace content | `hx-swap="innerHTML"` (default) |

---

## 💡 HTMX vs React

| Aspect | HTMX | React |
|--------|------|-------|
| **Learning Curve** | Rendah (HTML attributes) | Tinggi (JSX, hooks, state) |
| **Bundle Size** | 14KB | ~40KB (React + ReactDOM) |
| **Build Process** | ❌ Tidak perlu | ✅ Wajib (Webpack/Vite) |
| **Server Load** | Lebih tinggi | Lebih rendah |
| **SEO** | ✅ Native support | Butuh SSR/SSG |
| **Time to Interactive** | Instant | Perlu hydration |

---

## 🎯 When to Use HTMX?

### ✅ Perfect For:
- CRUD applications
- Admin dashboards
- Content-heavy websites
- Forms & validation
- Prototyping cepat
- Team backend-heavy

### ❌ Not Ideal For:
- Offline-first apps
- Real-time collaboration (Google Docs style)
- Heavy client-side logic
- Complex animations
- Mobile apps

---

## 🛠️ Next Steps

Mau explore lebih jauh? Coba:

- [ ] Tambah pagination
- [ ] Implementasi filter by category
- [ ] Tambah sorting (A-Z, Z-A)
- [ ] Integrasi database (PostgreSQL/MySQL)
- [ ] Deploy ke production (Railway/Fly.io)
- [ ] Tambah authentication

---

## 📚 Resources

- [HTMX Documentation](https://htmx.org/docs/)
- [HTMX Examples](https://htmx.org/examples/)
- [Go net/http Guide](https://pkg.go.dev/net/http)
- [HTML Templates in Go](https://pkg.go.dev/html/template)

---

## 🤝 Contributing

Contributions are welcome! Feel free to:

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

## 📝 License

This project is open source and available under the [MIT License](LICENSE).

---

<div align="center">

**⭐ Star this repo if you find it helpful!**

Made with ❤️ and Go

</div>