/*
 * Ent Schema API
 *
 * This is an auto generated API description made out of an Ent schema definition
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type EntriesBody struct {

	Content string `json:"content"`

	CreatedAt time.Time `json:"created_at"`
}