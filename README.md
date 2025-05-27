structure my project :

my-app/
│
├── backend/                  # Golang backend
│   ├── main.go               # نقطه شروع برنامه
│   ├── go.mod / go.sum       # مدیریت وابستگی‌ها
│   ├── config/               # تنظیمات محیطی و فایل‌های کانفیگ
│   ├── controllers/          # کنترلرها برای مدیریت درخواست‌ها
│   ├── models/               # مدل‌های دیتابیس
│   ├── routes/               # تعریف مسیرهای API
│   ├── services/             # منطق بیزینسی
│   └── utils/                # ابزارهای کمکی
│
├── frontend/                 # React frontend
│   ├── public/
│   ├── src/
│   │   ├── assets/           # تصاویر و فایل‌های استاتیک
│   │   ├── components/       # کامپوننت‌های قابل استفاده مجدد
│   │   ├── pages/            # صفحات مختلف (route-based)
│   │   ├── services/         # فراخوانی API‌ها
│   │   ├── App.jsx
│   │   └── main.jsx
│   └── package.json          # وابستگی‌های فرانت‌اند
│
├── .gitignore
├── README.md
└── docker-compose.yml        # اجرای همزمان Go و React
