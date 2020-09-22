package authorization

import (
	"apiservice/connections"
	"apiservice/models"
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"time"
)

func DeleteAuth(givenUuid string) error {

	db := connections.DBConn()
	tokenUser := models.TokenUser{}

	err := db.Where("access_uid = ?", givenUuid).Delete(&tokenUser).Error
	if err != nil {
		return err
	}
	return nil
}

func FetchAuth(authD *AccessDetails) (int64, error) {
	db := connections.DBConn()
	tokenUser := models.TokenUser{}

	err := db.Where("access_uid = ?", authD.AccessUuid).First(&tokenUser).Error
	// userid, err := client.Get(authD.AccessUuid).Result()
	if err != nil {
		return 0, err
	}
	// userID, _ := strconv.ParseUint(userid, 10, 64)
	return tokenUser.UserID, nil
}

func CreateAuth(userid uint64, td *TokenDetails) error {
	db := connections.DBConn()
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)

	tokenUser := models.TokenUser{UserID: int64(userid), AccessUid: td.AccessUuid, AtExpires: at}

	errAccess := db.Create(&tokenUser).Error
	if errAccess != nil {
		return errAccess
	}
	return nil
}

func CheckAuth(r *http.Request) error {
	tokenAuth, err := ExtractTokenMetadata(r)
	if err != nil {
		return err
	}

	_, err = FetchAuth(tokenAuth)
	if err != nil {
		return err
	}

	return nil
}

// func GetUserIDToken(r *http.Request) int64 {
// 	tokenAuth, err := ExtractTokenMetadata(r)
// 	if err != nil {
// 		return 0
// 	}
// 	return int64(tokenAuth.UserId)
// }

func CheckPasswordHash(password, encrytedPassword, salt string) bool {
	p := sha1.New()
	p.Write([]byte(password + salt))
	passwordHash := hex.EncodeToString(p.Sum(nil))
	if passwordHash == encrytedPassword {
		return true
	}
	return false
}
