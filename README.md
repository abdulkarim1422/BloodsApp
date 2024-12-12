# Blood Donation App
**`اللغة العربية في الأسفل`**

## Overview
The Blood Donation App is a web application designed to facilitate the process of blood donation. It allows users to register as donors, find patients in need of blood, and manage donations efficiently.

## Features
- **Donor Registration**: Users can register as blood donors by providing their personal details and blood type.
- **Patient Registration**: Users can register patients who need blood donations, specifying the required blood type and urgency.
- **Donation Management**: Admins can manage donations, matching donors with patients based on blood type compatibility.
- **Feedback System**: Send feedback to donors and patients via email.

## Technologies Used
- **Frontend**: HTML, CSS, JavaScript
- **Backend**: Go (Golang), Gin framework
- **Database**: PostgreSQL
- **ORM**: GORM

## Installation
1. Clone the repository:
    ```sh
    git clone https://github.com/abdulkarim1422/BloodsApp.git
    ```
2. Navigate to the project directory:
    ```sh
    cd BloodsApp
    ```
3. Install dependencies:
    ```sh
    go mod tidy
    ```
4. Configure the database connection in the `.env` file:
    ```env
    DB_USER=username
    DB_PASSWORD=password
    DB_HOST=localhost
    DB_PORT=5432
    DB_NAME=db_name
    ```
5. Run the application:
    ```sh
    go run main.go
    ```

## Usage
1. Open your web browser and navigate to `http://localhost:8080`.
2. Use the navigation menu to access different features of the application.
3. Register as a donor or patient, and manage donations through the admin dashboard.

## Contributing
Contributions are welcome! Please fork the repository and create a pull request with your changes.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

# تطبيق التبرع بالدم

## نظرة عامة
تطبيق التبرع بالدم هو تطبيق ويب مصمم لتسهيل عملية التبرع بالدم. يسمح للمستخدمين بالتسجيل كمتبرعين، والعثور على المرضى الذين يحتاجون إلى الدم، وإدارة التبرعات بكفاءة.

## الميزات
- **تسجيل المتبرعين**: يمكن للمستخدمين التسجيل كمتبرعين بالدم من خلال تقديم تفاصيلهم الشخصية وفصيلة الدم.
- **تسجيل المرضى**: يمكن للمستخدمين تسجيل المرضى الذين يحتاجون إلى تبرعات بالدم، مع تحديد فصيلة الدم المطلوبة ومدى الحاجة.
- **إدارة التبرعات**: يمكن للمسؤولين إدارة التبرعات، ومطابقة المتبرعين مع المرضى بناءً على توافق فصيلة الدم.
- **نظام التغذية الراجعة**: إرسال التغذية الراجعة إلى المتبرعين والمرضى عبر البريد الإلكتروني.

## التقنيات المستخدمة
- **الواجهة الأمامية**: HTML, CSS, JavaScript
- **الواجهة الخلفية**: Go (Golang), إطار عمل Gin
- **قاعدة البيانات**: PostgreSQL
- **ORM**: GORM

## التثبيت
1. استنساخ المستودع:
    ```sh
    git clone https://github.com/abdulkarim1422/BloodsApp.git
    ```
2. الانتقال إلى دليل المشروع:
    ```sh
    cd BloodsApp
    ```
3. تثبيت التبعيات:
    ```sh
    go mod tidy
    ```
4. تكوين اتصال قاعدة البيانات في ملف `.env`:
    ```env
    DB_USER=username
    DB_PASSWORD=password
    DB_HOST=localhost
    DB_PORT=5432
    DB_NAME=db_name
    ```
5. تشغيل التطبيق:
    ```sh
    go run main.go
    ```

## الاستخدام
1. افتح متصفح الويب الخاص بك وانتقل إلى `http://localhost:8080`.
2. استخدم قائمة التنقل للوصول إلى الميزات المختلفة للتطبيق.
3. سجل كمتبرع أو مريض، وقم بإدارة التبرعات من خلال لوحة التحكم الخاصة بالمسؤول.

## المساهمة
المساهمات مرحب بها! يرجى استنساخ المستودع وإنشاء طلب سحب مع التغييرات الخاصة بك.

## الترخيص
هذا المشروع مرخص بموجب ترخيص MIT. راجع ملف [LICENSE](LICENSE) لمزيد من التفاصيل.