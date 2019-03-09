package seeds

import (
	"github.com/icrowley/fake"
	"github.com/jinzhu/gorm"
	"github.com/melardev/gogonic_gorm_api_crud/models"
	"math/rand"
	"time"
)

func randomInt(min, max int) int {

	return rand.Intn(max-min) + min
}
func Seed(db *gorm.DB) {
	fake.Seed(time.Now().Unix())
	var countTodos int
	db.Model(&models.Todo{}).Count(&countTodos)
	todosToSeed := 12
	todosToSeed -= countTodos

	for i := 0; i < todosToSeed; i++ {
		completed := true
		if randomInt(0, 20)%2 == 0 {
			completed = false
		}
		db.Create(&models.Todo{
			Title:       fake.WordsN(randomInt(2, 4)),
			Description: fake.SentencesN(randomInt(1, 3)),
			Completed:   completed,
		})
	}
}
