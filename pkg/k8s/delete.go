package k8s

// // ParseDeleteResource doc
// func ParseDeleteResource(c *fiber.Ctx, kc KClient, resource string) error {
// 	p := new(model.UnstructuredDeleteRequest)
// 	if err := c.BodyParser(&p); err != nil {
// 		return errors.New("Error parsing request payload with error: " + err.Error())
// 	}

// 	if !strings.Contains(p.TestLabel, "stress-test-label=") {
// 		return errors.New("Error: test_label must start with \"stress-test-label\"")
// 	}
// 	return nil

// 	listOption := metav1.ListOptions{
// 		LabelSelector: p.TestLabel,
// 	}
// 	kc.DynamicClient.Resource(resource).Namespace(namespace).Delete(kc.Ctx, listOption)
// }
