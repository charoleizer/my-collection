package domain

import (
	"github.com/charoleizer/my-collection/src/models"
)

// Checks if the figure's name is Luffy
func IsLuffy(figure models.Figures) bool {
	return figure.Name == "Luffy"
}
