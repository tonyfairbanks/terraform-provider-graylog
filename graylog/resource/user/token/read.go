package token

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/tonyfairbanks/terraform-provider-graylog/graylog/client"
	"github.com/tonyfairbanks/terraform-provider-graylog/graylog/util"
)

func read(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}
	data, resp, err := cl.UserToken.Get(ctx, d.Get("user_id").(string), d.Id(), cl.APIVersion)
	if err != nil {
		return util.HandleGetResourceError(d, resp, err)
	}
	return setDataToResourceData(d, data, cl.APIVersion)

}
