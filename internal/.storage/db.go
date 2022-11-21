package storage

/*
CREATE TABLE users
(
	id SERIAL,
	username varchar(32) NOT NULL DEFAULT '',
	user_guid varchar(256) NOT NULL DEFAULT '',
	user_hash varchar(256) NOT NULL DEFAULT ''
);        //ENGINE=InnoDB DEFAULT CHARSET=latin1;
*/

/*
CREATE TABLE cookie
(
	user_guid_cookie varchar(256) NOT NULL DEFAULT '',
	cookie varchar(256) NOT NULL DEFAULT '',
	FOREIGN KEY (user_guid_cookie) REFERENCES users (user_guid)
);
*/

// func Connect() *pgx.Conn {
// 	urlExample := "postgres://postgres:password@localhost:5432/gophermarket" //os.Getenv("DATABASE_URL")
// 	conn, err := pgx.Connect(context.Background(), urlExample)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// 		return nil
// 	}
// 	fmt.Println("connection is established")
// 	return conn
// 	// var name int
// 	// var email string
// 	// defer conn.Close(context.Background())
// 	// err = conn.QueryRow(context.Background(), "select * from customers").Scan(&name, &email)
// 	// if err != nil {
// 	// 	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
// 	// 	os.Exit(1)
// 	// }

// 	// fmt.Println(name, email)
// }

// func GetGUID(conn *pgx.Conn, username string) (string, error) {
// 	var guid string
// 	if err := conn.QueryRow(context.Background(), "select user_guid from users where username=$1", username).Scan(&guid); err != nil {
// 		return "", err
// 	}
// 	return guid, nil
// }

// func AddUser( /*ctx context.Context,*/ conn *pgx.Conn, userInfo info.User) (bool, error) {
// 	var guid string
// 	ctx := context.Background()
// 	if err := conn.QueryRow(context.Background(), "select user_guid from users where username=$1", userInfo.Login).Scan(&guid); err != nil {
// 		if errors.Is(err, pgx.ErrNoRows) {
// 			uuid.Must(uuid.NewRandom())
// 			guid := fmt.Sprintf("%s", uuid.New())
// 			// guid = String(uuid.New())
// 			hash := service.MyHash{}
// 			hashString := hash.GenerateHash(userInfo.Password)
// 			fmt.Println(guid, hashString)
// 			_, err := conn.Exec(ctx, "INSERT INTO users(username, user_guid, user_hash) VALUES($1,$2,$3)", userInfo.Login, guid, hashString)
// 			if err != nil {
// 				return true, err
// 			}
// 			return true, nil
// 		} else {
// 			return false, err
// 		}
// 	}
// 	return false, nil
// }

// func GenCook(username string) (string, error) {
// 	conn := Connect()
// 	hash := service.MyHash{}
// 	guid, err := GetGUID(conn, username)
// 	if err != nil {
// 		return "", err
// 	}
// 	date := time.Now()
// 	formatDate := date.Format("20060102150405")
// 	seq := guid + formatDate
// 	cook := hash.GenerateHash(seq)
// 	return cook, nil
// }
