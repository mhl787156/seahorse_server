package db

import "github.com/mhl787156/seahorse_server/models"

func (c *DbConn) WriteSession(session models.Session) error {
	// resp, err := SessionTable.Insert(session).RunWrite(c.Session)
	// if err != nil {
	// 	return err
	// } else if resp.Errors != 0 {
	// 	return errors.New("Database insert failed")
	// } else if resp.Inserted != 1 {
	// 	return errors.New("Incorrect number of sessions inserted")
	// } else {
	// 	return nil
	// }
	return nil
}

func (c *DbConn) GetProject(pid string) (*models.Project, bool, error) {
	// resp, err := ProjectTable.Get(pid).Run(c.Session)
	// if err != nil {
	// 	return nil, false, err
	// }
	// defer resp.Close()

	// //Check that a result was found
	// if resp.IsNil() {
	// 	//No results were found
	// 	return nil, false, nil
	// } else {
	// 	session := models.Project{}
	// 	err = resp.One(&session)
	// 	if err != nil {
	// 		return nil, false, err
	// 	} else {
	// 		return &session, true, nil
	// 	}
	// }
	return nil, false, nil
}

func (c *DbConn) GetProjectContent(pid string) (*models.ProjectContent, error) {
	// resp, err := ProjectCTable.Get(pid).Run(c.Session)
	// if err != nil {
	// 	return nil, err
	// }
	// defer resp.Close()

	// //Check that a result was found
	// if resp.IsNil() {
	// 	//No results were found
	// 	return nil, errors.New("wat")
	// } else {
	// 	session := models.ProjectContent{}
	// 	err = resp.One(&session)
	// 	if err != nil {
	// 		return nil, err
	// 	} else {
	// 		return &session, nil
	// 	}
	// }
	return nil, nil
}

func (c *DbConn) WriteProject(me *models.User) (*models.Project, error) {
	// proj := models.NewDefaultProject(me.Id)
	// proj.Id = c.GetUUID()
	// projContent := models.NewDefaultProjectContent()
	// projContent.Id = proj.Id

	// me.Projects = append(me.Projects, proj.Id)
	// found, err := c.ModifyUser(me)
	// if err != nil {
	// 	return nil, err
	// } else if !found {
	// 	return nil, errors.New("Could not find that user")
	// }

	// resp, err := ProjectTable.Insert(proj).RunWrite(c.Session)
	// _, _ = ProjectCTable.Insert(projContent).RunWrite(c.Session)
	// if err != nil {
	// 	return nil, err
	// } else if resp.Errors != 0 {
	// 	return nil, errors.New("Database insert failed")
	// } else if resp.Inserted != 1 {
	// 	return nil, errors.New("Incorrect number of projects inserted")
	// } else {
	// 	return &proj, nil
	// }
	return nil, nil
}

func (c *DbConn) ModifyProjectContent(proj *models.ProjectContent) (bool, error) {
	// res, err := ProjectCTable.Get(proj.Id).Update(*proj).RunWrite(c.Session)
	// if err != nil {
	// 	return false, err
	// } else if res.Replaced == 0 {
	// 	return false, nil
	// } else {
	// 	return true, err
	// }
	return false, nil
}

func (c *DbConn) ModifyProject(proj *models.Project) (bool, error) {
	// res, err := ProjectTable.Get(proj.Id).Update(*proj).RunWrite(c.Session)
	// if err != nil {
	// 	return false, err
	// } else if res.Replaced == 0 {
	// 	return false, nil
	// } else {
	// 	return true, err
	// }
	return false, nil
}

func (c *DbConn) GetSession(sessionKey string) (*models.Session, bool, error) {
	// query, err := SessionTable.Get(sessionKey).Run(c.Session)
	// if err != nil {
	// 	fmt.Printf("Could not find session for key %s\n", sessionKey)
	// 	return nil, false, err
	// }

	// defer query.Close()

	// //Check that a result was found
	// if query.IsNil() {
	// 	//No results were found
	// 	fmt.Printf("Could not find session for key %s\n", sessionKey)
	// 	return nil, false, nil
	// } else {
	// 	foundSession := models.Session{}
	// 	err = query.One(&foundSession)
	// 	if err != nil {
	// 		return nil, false, err
	// 	} else {
	// 		return &foundSession, true, nil
	// 	}
	// }
	return nil, false, nil
}

func (c *DbConn) GetSnippets(uid string) ([]string, error) {
	// query, err := SnippetTable.Filter(map[string]interface{}{
	// 	"owner": uid,
	// }).Field("id").Run(c.Session)
	// if err != nil {
	// 	return nil, err
	// }

	// defer query.Close()

	// //Check that a result was found
	// if query.IsNil() {
	// 	//No results were found
	// 	return nil, nil
	// } else {
	// 	var foundSnippets []string
	// 	err = query.All(&foundSnippets)
	// 	if err != nil {
	// 		return nil, err
	// 	} else {
	// 		return foundSnippets, nil
	// 	}
	// }
	return nil, nil
}

func (c *DbConn) GetFBUser(fid string) (*models.User, bool, error) {
	// fmt.Printf("Authenticating FID %s\n", fid)
	// query, err := UserTable.Filter(map[string]interface{}{"fid": fid}).Run(c.Session)
	// if err != nil {
	// 	fmt.Println("Failed to search db")
	// 	return nil, false, err
	// }

	// defer query.Close()

	// //Check that a result was found
	// if query.IsNil() {
	// 	fmt.Println("Failed to find user")
	// 	//No results were found
	// 	return nil, false, nil
	// } else {
	// 	foundUser := models.User{}
	// 	err = query.One(&foundUser)
	// 	if err != nil {
	// 		fmt.Printf("There was an error: %s\n", err)
	// 		return nil, false, err
	// 	} else {
	// 		return &foundUser, true, nil
	// 	}
	// }
	return nil, false, nil
}

func (c *DbConn) RemoveUserFromProject(pid string, uid string) (bool, error) {
	// res, err := UserTable.Get(uid).Update(map[string]interface{}{
	// 	"projects": gorethink.Row.Field("projects").SetDifference([]string{pid}),
	// }).RunWrite(c.Session)
	// if err != nil {
	// 	return false, err
	// } else if res.Replaced == 0 {
	// 	return false, nil
	// } else {
	// 	res, err := ProjectTable.Get(pid).Update(map[string]interface{}{
	// 		"owner": gorethink.Row.Field("owner").SetDifference([]string{uid}),
	// 	}).RunWrite(c.Session)
	// 	if err != nil {
	// 		return false, err
	// 	} else if res.Replaced == 0 {
	// 		return false, nil
	// 	} else {
	// 		res, err := ProjectTable.Get(pid).Field("owner").Count().Run(c.Session)
	// 		defer res.Close()
	// 		if err != nil {
	// 			return false, err
	// 		} else {
	// 			var cnt int
	// 			res.One(&cnt)
	// 			if cnt == 0 {
	// 				_, _ = ProjectTable.Get(pid).Delete().RunWrite(c.Session)
	// 			}
	// 			return true, nil
	// 		}
	// 	}
	// }
	return false, nil
}

func (c *DbConn) GetSnippet(sid string) (*models.Snippet, bool, error) {
	// resp, err := SnippetTable.Get(sid).Run(c.Session)
	// if err != nil {
	// 	return nil, false, err
	// }
	// defer resp.Close()

	// //Check that a result was found
	// if resp.IsNil() {
	// 	//No results were found
	// 	return nil, false, nil
	// } else {
	// 	session := models.Snippet{}
	// 	err = resp.One(&session)
	// 	if err != nil {
	// 		return nil, false, err
	// 	} else {
	// 		return &session, true, nil
	// 	}
	// }
	return nil, false, nil
}

func (c *DbConn) GetSnippetContent(sid string) (*models.SnippetContent, error) {
	// resp, err := SnippetCTable.Get(sid).Run(c.Session)
	// if err != nil {
	// 	return nil, err
	// }
	// defer resp.Close()

	// //Check that a result was found
	// if resp.IsNil() {
	// 	//No results were found
	// 	return nil, errors.New("wat")
	// } else {
	// 	session := models.SnippetContent{}
	// 	err = resp.One(&session)
	// 	if err != nil {
	// 		return nil, err
	// 	} else {
	// 		return &session, nil
	// 	}
	// }
	return nil, nil
}

func (c *DbConn) WriteSnippet(me *models.User) (*models.Snippet, error) {
	// snippet := models.NewDefaultSnippet(me.Id)
	// snippet.Id = c.GetUUID()
	// snippetContent := models.SnippetContent{}
	// snippetContent.Id = snippet.Id

	// resp, err := SnippetTable.Insert(snippet).RunWrite(c.Session)
	// _, _ = SnippetCTable.Insert(snippetContent).RunWrite(c.Session)
	// if err != nil {
	// 	return nil, err
	// } else if resp.Errors != 0 {
	// 	return nil, errors.New("Database insert failed")
	// } else if resp.Inserted != 1 {
	// 	return nil, errors.New("Incorrect number of snippets inserted")
	// } else {
	// 	return &snippet, nil
	// }
	return nil, nil
}

func (c *DbConn) ModifySnippetContent(snippet *models.SnippetContent) (bool, error) {
	// res, err := SnippetCTable.Get(snippet.Id).Update(*snippet).RunWrite(c.Session)
	// if err != nil {
	// 	return false, err
	// } else if res.Replaced == 0 {
	// 	return false, nil
	// } else {
	// 	return true, err
	// }
	return false, nil
}

func (c *DbConn) ModifySnippet(snippet *models.Snippet) (bool, error) {
	// res, err := SnippetTable.Get(snippet.Id).Update(*snippet).RunWrite(c.Session)
	// if err != nil {
	// 	return false, err
	// } else if res.Replaced == 0 {
	// 	return false, nil
	// } else {
	// 	return true, err
	// }
	return false, nil
}

func (c *DbConn) DeleteSnippet(sid string) error {
	// resp, err := SnippetTable.Get(sid).Delete().RunWrite(c.Session)
	// if err != nil {
	// 	return err
	// } else if resp.Errors != 0 {
	// 	return errors.New("Database update failed")
	// } else if resp.Inserted != 1 {
	// 	return errors.New("Could not find that snippet")
	// } else {
	// 	_, _ = SnippetCTable.Get(sid).Delete().RunWrite(c.Session)
	// 	return nil
	// }
	return nil
}
