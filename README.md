# Booings and Reservations Project.

- Built in Go version 1.23.3
- Uses the [chi router](github.com/alexedwards/scs/v2)
- Uses [alex edwards SCS](github.com/go-chi/chi/v5) seassion management
- Uses [nsurf](github.com/justinas/nosurf)

## **Project Overview**
The **Booings and Reservations System** is a full-stack web application designed to simplify the booking process for customers and streamline reservation management for hotels. This project provides a seamless user experience, integrating a responsive front-end with a robust back-end.

---

## **Setup and Installation**

### **Prerequisites:**
- Go installed on your system.
- PostgreSQL database set up and running.
- Optional: MailHog installed for email testing.

### **Steps to Run Locally:**
1. Clone the repository:
   ```bash
   git clone https://github.com/fuad7161/Golang
   cd Golang/cmd/web
   ```

2. Set up the environment variables for database connection:
- Update your PostgreSQL connection string in the application config.

3. Install required Go packages:
   ```bash
   go mod tidy
   ```

4. Run the application:
   ```bash
   go run .
   ```

5. Access the application in your browser at `http://localhost:8080`.

---
## **Technology and Tools**

### **Front-end:**
- **HTML**, **CSS**, **JavaScript**, and **Bootstrap** for building a responsive and user-friendly interface.

### **Back-end:**
- **Golang** for server-side development.
- Key Libraries/Packages:
  - **Routing and Middleware:** `github.com/go-chi/chi`, `github.com/gorilla/mux`
  - **Session Management:** `github.com/alexedwards/scs/v2`
  - **Validation:** `github.com/asaskevich/govalidator`
  - **CSRF Protection:** `github.com/justinas/nosurf`
  - **Database Handling:** `github.com/jackc/pgx/v4`
  - **Email Services:** `github.com/xhit/go-simple-mail/v2`, `github.com/mailhog`

### **Database:**
- **PostgreSQL** for efficient data storage and management.
  1. **Room_Restrictions** - Stores room-specific restrictions (e.g., availability, maintenance).
  2. **Room_reservations** - Tracks the reservations made for rooms.
  3. **Reservations** - Stores the overall reservation details.
  4. **Restrictions** - General restrictions (e.g., date, booking limits).
  5. **Room** - Contains information about each room (e.g., type, capacity, price).
  6. **User** - Stores user information (e.g., name, email, password).
  
### **Testing and Mail Services:**
- **MailHog** for email testing.
- **Static File Serving:** `http.FileServer` for serving static assets.
- **Postman**  for API validation.

---

## **Future Enhancements**
- Implement user registration functionality.
- Add payment gateway integration.
- Enhance the admin dashboard with analytics and reports.
- Improve UI/UX with advanced front-end frameworks like Vue.js or React.

---

## **End Points**
### **Public APIs:**
- **Home and Information Pages:**
    - `/` - Home page
    - `/about` - About us
    - `/contact` - Contact page

- **User Authentication:**
    - `/user/login` (GET, POST) - Login functionality
    - `/user/logout` - Logout endpoint

- **Room Reservations:**
    - `/reservation` (GET, POST) - Make a reservation
    - `/reservation-summary` - View reservation summary
    - `/search-availability` (GET, POST) - Search room availability
    - `/choose-room/{id}` - Choose a specific room
    - `/book-room` - Book a selected room

- **Static File Handling:**
    - `/static/*` - Serves static files (CSS, JS, images)

### **Admin APIs:**
- **Dashboard and Management:**
    - `/admin/dashboard` - Admin dashboard
    - `/admin/reservations-new` - View new reservations
    - `/admin/reservations-all` - View all reservations
    - `/admin/reservations-calendar` (GET, POST) - Manage reservations via calendar
    - `/admin/reservations/{src}/{id}` (GET, POST) - View and edit individual reservations

### **Database Management APIs:**
- `/addUser` (GET, POST) - Add a new user
- `/show-all-User` - Display all users

---

---

## **Contact**
For any inquiries or feedback, feel free to reach out:
- **Email:** fuadul202@gmail.com
- **GitHub:** https://github.com/fuad7161
