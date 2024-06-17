package main

func commandViewProgress(c *config) error {
	c.printBooks() // add progress e.g: pages 291/401
	return nil
}