package rudderstack

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTransformation() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTransformationCreate,
		ReadContext:   resourceTransformationRead,
		UpdateContext: resourceTransformationUpdate,
		DeleteContext: resourceTransformationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the transformation.",
			},
			"code": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The JavaScript code for the transformation.",
			},
			"is_published": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether the transformation is published.",
			},
			"test_json": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The JSON used to test the transformation.",
			},
		},
	}
}

func resourceTransformationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	transformation := &Transformation{
		Name:        d.Get("name").(string),
		Code:        d.Get("code").(string),
		IsPublished: d.Get("is_published").(bool),
	}

	if v, ok := d.GetOk("test_json"); ok {
		transformation.TestJSON = v.(string)
	}

	res, err := c.Transformations.Create(ctx, transformation)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(res.ID)

	return resourceTransformationRead(ctx, d, m)
}

func resourceTransformationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	transformation, err := c.Transformations.Get(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", transformation.Name)
	d.Set("code", transformation.Code)
	d.Set("is_published", transformation.IsPublished)
	d.Set("test_json", transformation.TestJSON)

	return nil
}

func resourceTransformationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	transformation := &Transformation{
		ID:          d.Id(),
		Name:        d.Get("name").(string),
		Code:        d.Get("code").(string),
		IsPublished: d.Get("is_published").(bool),
	}

	if v, ok := d.GetOk("test_json"); ok {
		transformation.TestJSON = v.(string)
	}

	_, err := c.Transformations.Update(ctx, transformation)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceTransformationRead(ctx, d, m)
}

func resourceTransformationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	err := c.Transformations.Delete(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
