package store

import (
	"code-runner/internal/models"
	"code-runner/internal/tools"
	"context"
	"github.com/docker/docker/pkg/testutil/assert"
	"testing"
)

func Test_Store(t *testing.T) {
	ctx := context.Background()
	store := InitMongoStore(ctx)
	store.TestConnection(ctx)
}

func Test_CreateWorkspace(t *testing.T) {
	ctx := context.Background()
	store := InitMongoStore(ctx)
	store.TestConnection(ctx)
	ws := &models.Workspace{
		Owner: tools.GenerateRandomString(10),
	}
	_, err := store.CreateWorkspace(ctx, ws)
	assert.Equal(t, err, nil)
	wsGet, err := store.GetWorkspace(ctx, ws.Owner)
	assert.Equal(t, ws.Owner, wsGet.Owner)
	assert.Equal(t, err, nil)
}

func Test_CreateApp(t *testing.T) {
	ctx := context.Background()
	store := InitMongoStore(ctx)
	store.TestConnection(ctx)
	owner := tools.GenerateRandomString(10)
	name := tools.GenerateRandomString(10)
	ws := &models.Workspace{
		Owner: owner,
	}
	_, err := store.CreateWorkspace(ctx, ws)
	assert.Equal(t, err, nil)
	wsGet, err := store.GetWorkspace(ctx, ws.Owner)
	assert.Equal(t, ws.Owner, wsGet.Owner)
	assert.Equal(t, err, nil)

	app := &models.App{
		Owner: ws.Owner,
		Name:  name,
	}

	_, err = store.CreateApp(ctx, app)
	assert.Equal(t, err, nil)
	appGet, err := store.GetApp(ctx, owner, app.Name)
	assert.NilError(t, err)
	assert.Equal(t, appGet.Name, name)
	assert.Equal(t, appGet.Status, models.INIT)
	err = store.DeleteApp(ctx, owner, name)
	assert.NilError(t, err)
	appGet, err = store.GetApp(ctx, owner, name)
	assert.NotNil(t, err)
}

func Test_ListApp(t *testing.T) {
	ctx := context.Background()
	store := InitMongoStore(ctx)
	store.TestConnection(ctx)
	owner := tools.GenerateRandomString(10)
	ws := &models.Workspace{
		Owner: owner,
	}
	_, err := store.CreateWorkspace(ctx, ws)
	assert.Equal(t, err, nil)
	wsGet, err := store.GetWorkspace(ctx, ws.Owner)
	assert.Equal(t, ws.Owner, wsGet.Owner)
	assert.Equal(t, err, nil)

	app := &models.App{
		Owner: ws.Owner,
		Name:  tools.GenerateRandomString(10),
	}
	app2 := &models.App{
		Owner: ws.Owner,
		Name:  tools.GenerateRandomString(10),
	}

	_, err = store.CreateApp(ctx, app)
	assert.Equal(t, err, nil)
	_, err = store.CreateApp(ctx, app2)
	assert.Equal(t, err, nil)

	apps, err := store.GetApps(ctx, owner)
	assert.Equal(t, len(apps), 2)

	wsWithApps, err := store.GetWorkspace(ctx, owner)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(wsWithApps.Apps), 2)
}

func Test_UpdateApp(t *testing.T) {
	ctx := context.Background()
	store := InitMongoStore(ctx)
	store.TestConnection(ctx)
	owner := tools.GenerateRandomString(10)
	name := tools.GenerateRandomString(10)
	url := tools.GenerateRandomString(10)

	ws := &models.Workspace{
		Owner: owner,
	}

	_, err := store.CreateWorkspace(ctx, ws)
	assert.Equal(t, err, nil)

	app := &models.App{
		Owner: owner,
		Name:  name,
	}

	_, err = store.CreateApp(ctx, app)
	assert.Equal(t, err, nil)
	app, err = store.GetApp(ctx, owner, name)
	assert.Equal(t, err, nil)
	assert.Equal(t, app.Name, name)
	assert.Equal(t, app.Url, "")
	app.Url = url
	updated, err := store.UpdateApp(ctx, app)
	assert.Equal(t, err, nil)
	assert.Equal(t, updated.Url, app.Url)
}

func Test_StatusApp(t *testing.T) {
	ctx := context.Background()
	store := InitMongoStore(ctx)
	store.TestConnection(ctx)

	owner := tools.GenerateRandomString(10)
	name := tools.GenerateRandomString(10)

	app := &models.App{
		Owner: owner,
		Name:  name,
	}

	_, err := store.CreateApp(ctx, app)
	assert.Equal(t, err, nil)
	app, err = store.GetApp(ctx, owner, name)
	assert.NilError(t, err)
	assert.Equal(t, app.Name, name)
	assert.Equal(t, app.Status, models.INIT)
	app.Status = models.BUILDING
	updated, err := store.UpdateApp(ctx, app)
	assert.NilError(t, err)
	assert.Equal(t, updated.Status, models.BUILDING)
	app.Status = models.READY
	updated, err = store.UpdateApp(ctx, app)
	assert.NilError(t, err)
	assert.Equal(t, updated.Status, models.READY)
	app.Status = models.READY
	updated, err = store.UpdateApp(ctx, app)
	assert.NilError(t, err)
	assert.Equal(t, updated.Status, models.READY)
	app.Status = models.RUNNING
	updated, err = store.UpdateApp(ctx, app)
	assert.NilError(t, err)
	assert.Equal(t, updated.Status, models.RUNNING)
}
