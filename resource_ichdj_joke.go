package main

import (
	"encoding/json"
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func resourceRandomJoke() *schema.Resource {
	return &schema.Resource{
		Create: resourceRandomJokeCreate,
		Read:   resourceRandomJokeRead,
		Update: resourceRandomJokeUpdate,
		Delete: resourceRandomJokeDelete,

		Schema: map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},

			"joke": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"status": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

type RespJsonDecoded struct {
	Id     string `json:"id"`
	Joke   string `json:"joke"`
	Status int64  `json:"status"`
}

func resourceRandomJokeCreate(d *schema.ResourceData, m interface{}) error {
	// Initialize new GET request.
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)

	if err != nil {
		log.Printf("[INFO] Could not initialize 'http.NewRequest' to 'https://icanhazdadjoke.com/'.")
		return err
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{
		Timeout: time.Second * 5,
	}

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Printf("[INFO] Could not read response body.")
		return err
	}

	var respJsonDecoded RespJsonDecoded
	// Convert respBody to format json.Unmarshal expects (bytes).
	err = json.Unmarshal([]byte(string(respBody)), &respJsonDecoded)

	if err != nil {
		log.Printf("[INFO] Could not unmarshal JSON response.")
		return err
	}

	d.SetId(respJsonDecoded.Id) // Set Terraform ID as joke ID.
	d.Set("joke", respJsonDecoded.Joke)
	d.Set("status", respJsonDecoded.Status)

	return nil
}

func resourceRandomJokeRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceRandomJokeUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceRandomJokeRead(d, m)
}

func resourceRandomJokeDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
