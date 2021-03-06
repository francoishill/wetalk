Important to CLONE first
========================
Do not edit this code directly, after installing it use the clone tool (clonebeegowebapp) as described in "Quick Start" below to get you up and running.


# What is this?

An open source project to create a skeleton of a [Beego](https://github.com/astaxie/beego) website.

The aim is to have all the basics of a website, like User registration, login, password reset, etc. It will also include clientside libraries like AngularJS and Bootstrap.


This package was initially forked from [wetalk](https://github.com/beego/wetalk) and stripped down to give a basic web application.


# Included in this package

The stripped down version includes the following features:

0. **Authentication** (register, activate email, reset password, login, XSRF, etc)
0. **Admin** side to manage the registered users
0. **Migration** (makes use [Beego](https://github.com/astaxie/beego)/orm syncdb) to do migration and create log files
0. **Localization** using [i18n](https://github.com/beego/i18n)
0. **Two config** files, one with general settings (like app name) and the other with machine specific settings (like runmode, http_port, etc)
0. **[Compression](https://github.com/beego/compress)** of static files (like css, js, etc)
0. **MVC** pattern is used which is part of Beego splits our code up into Routers (controllers), Models and Views
0. We make use of a Master view (in views/master.html) to define the default layout of the site, ``{{.LayoutContent}}`` defines where the other templates are rendered
0. Making use of [Beego](https://github.com/astaxie/beego) also gives us the ORM capability


We have also included the following libraries into the design (they are laid out in the conf/compress.json file):

0. [AngularJS](http://angularjs.org/), created by Google and is the missing link between HTML and Javascript
0. [Bootstrap](http://getbootstrap.com/2.3.2/), created by Twitter and is a front-end framework
0. [jQuery](http://jquery.com/), do we really need to say what jQuery is?
0. [SASS](http://sass-lang.com/), a CSS-precompiler


# Installation

0. #### Dependencies to be installed

    ```bash
    go get github.com/astaxie/beego
    go get github.com/astaxie/bee
    go get github.com/Unknwon/goconfig
    go get github.com/beego/i18n
    go get github.com/howeyc/fsnotify
    go get github.com/beego/compress
    go get github.com/go-sql-driver/mysql
    go get github.com/francoishill/runsass
    go get github.com/francoishill/beegowebapp
    ```

0. #### Now install the cloning tool

    ```bash
    cd $GOPATH/src/github.com/francoishill/beegowebapp
    ./install_cloner.sh
    ```

0. #### Additional (non-default) packages that can be used by this application/beego:

    0. memcache: https://github.com/youtube/vitess
    0. redis: github.com/garyburd/redigo/redis
    0. x2jXML: github.com/clbanning/x2j
    0. goyaml2: github.com/wendal/goyaml2
    0. postgres: github.com/lib/pq
    0. sqlite3: github.com/mattn/go-sqlite3
    0. websockets: github.com/garyburd/go-websocket/websocket


# Quick Start

0. #### Generate your first web app (based on the [Beego](https://github.com/astaxie/beego) framework)

    ```bash
    cd /temp
    clonebeegowebapp blank --foldername="myfirstapp"
    ```
    
0. #### Important changes required in file `app_general.ini` found in the `conf` folder.

    ```ini
    [app]
    app_name = My Web Application
    
    author_firstname = Francois
    author_surname = Hill

    session_path = sess_db_username:sess_db_password@/sess_db_name?charset=utf8
    session_name = myapp_session
    mail_user = Myapp Mailer
    mail_from = admin@myapp.com
    secret_key = 1234abc890def567
    
    [orm]
    data_source = main_db_username:main_db_password@/main_db_name?charset=utf8
    ```
    
0. #### Have a look at the file `app_machine_specific.ini` too.


0. #### Change the following in localization files (`conf/locale_....ini`)

    ```ini
    AF =
    EN-US =
    app_name =
    app_intro =
    app_desc =
    app_keywords =
    app_welcome_message =
    app_brand =
    app_copyright =
    slogan =
    app_meta_author_firstname =
    app_meta_author_surname =
    ```
    
    Note: the values of `AF` and `EN-US` are what will be displayed to the user when choosing languages. To create another language just add a new `.ini` file in the `conf` folder with its localization key and add its display value to each localization `.ini` file.
    
    For instance to add chinese create a file named `locale_zh-CN.ini`. Then add `zh-CN = Chinese` to each localization `.ini` file, including `locale_zh-CN.ini` itsself.
    
    
0. #### If you configured mysql to be the session provider, create its required table with the following SQL:

    ```sql
    CREATE TABLE `session` (
        `session_key` char(64) NOT NULL,
        `session_data` blob,
        `session_expiry` int(11) unsigned NOT NULL,
        PRIMARY KEY (`session_key`)
    ) ENGINE=MyISAM DEFAULT CHARSET=utf8;
    ```
    
0. #### Run the migration to generate the required tables.

    Run the script `db/migrate_windows.bat` (or `db/migrate_linux.sh` if on linux) to create all the tables for models registerd with `orm.RegisterModel`. By default this will be like the *User* model.
    
    This migration status will be logged into a timestamped file inside *db/migrations/windows* or *db/migrations/linux*.
    
    **Double check** this log to ensure everything ran smoothly.
    
    
## Running on another machine

To run as a stand-alone web server on another machine you will only need **1. a database** and **2. to copy the following files/folders**:

0. conf
0. static
0. static_source
0. views
0. ??.exe (your applications' executable, on windows it will have the extension .exe)


#### Portable MySql tools (windows only)

- A portable MySql server: [https://github.com/francoishill/windows_portable_mysql](https://github.com/francoishill/windows_portable_mysql)
- A portable Admin tool for MySql: [https://github.com/francoishill/windows_heidisql_portable](https://github.com/francoishill/windows_heidisql_portable)




# Contribution

Please feel free to contribute or to contact me.

## License

[The MIT License (MIT)](http://opensource.org/licenses/MIT).

