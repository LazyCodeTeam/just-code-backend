package usecase

import (
	"context"
	"reflect"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/port"
)

type UploadContent struct {
	transactionFactory port.TransactionFactory
}

func NewUploadContent(transactionFactory port.TransactionFactory) *UploadContent {
	return &UploadContent{
		transactionFactory: transactionFactory,
	}
}

func (u *UploadContent) Invoke(
	ctx context.Context,
	params []model.ExpectedTechnology,
	dryRun bool,
) error {
	transaction, err := u.transactionFactory.Begin(ctx)
	if err != nil {
		return err
	}
	defer transaction.Rollback(ctx)

	repo := transaction.ContentRepository(ctx)
	technologies, sections, tasks, err := getCurrentContent(ctx, repo)
	if err != nil {
		return err
	}
	expectedTechnologies, expectedSections, expectedTasks := getExpectedContent(params)
	err = alignStateToExpected(
		ctx,
		expectedTechnologies,
		technologies,
		repo.UpsertTechnology,
		repo.DeleteTechnologyById,
		func(technology model.Technology) string { return technology.Id },
	)
	if err != nil {
		return err
	}
	err = alignStateToExpected(
		ctx,
		expectedSections,
		sections,
		repo.UpsertSection,
		repo.DeleteSectionById,
		func(section model.Section) string { return section.Id },
	)
	if err != nil {
		return err
	}
	err = alignStateToExpected(
		ctx,
		expectedTasks,
		tasks,
		repo.UpsertTask,
		repo.DeleteTaskById,
		func(task model.Task) string { return task.Id },
	)
	if err != nil {
		return err
	}

	if dryRun {
		return nil
	}

	err = transaction.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func alignStateToExpected[T any](
	ctx context.Context,
	expected []T,
	actual map[string]T,
	upsertItem func(context.Context, T) error,
	deleteItem func(context.Context, string) error,
	idResolver func(T) string,
) error {
	for _, expectedItem := range expected {
		id := idResolver(expectedItem)
		actualItem, ok := actual[id]
		if ok && reflect.DeepEqual(expectedItem, actualItem) {
			delete(actual, id)
			continue
		}
		err := upsertItem(ctx, expectedItem)
		if err != nil {
			return err
		}
		delete(actual, id)
	}
	for _, actualItem := range actual {
		id := idResolver(actualItem)
		err := deleteItem(ctx, id)
		if err != nil {
			return err
		}
	}
	return nil
}

func getExpectedContent(
	expectedTechnologies []model.ExpectedTechnology,
) ([]model.Technology, []model.Section, []model.Task) {
	technologies := make([]model.Technology, 0, len(expectedTechnologies))
	sections := make([]model.Section, 0)
	tasks := make([]model.Task, 0)

	for _, expectedTechnology := range expectedTechnologies {
		technologies = append(technologies, expectedTechnology.Technology)
		for _, expectedSection := range expectedTechnology.ExpectedSections {
			sections = append(sections, expectedSection.Section)
			for _, expectedTask := range expectedSection.ExpectedTasks {
				tasks = append(tasks, expectedTask)
			}
		}
	}
	return technologies, sections, tasks
}

func getCurrentContent(
	ctx context.Context,
	repo port.ContentRepository,
) (map[string]model.Technology, map[string]model.Section, map[string]model.Task, error) {
	technologies, err := getAllTechnologiesAsMap(ctx, repo)
	if err != nil {
		return nil, nil, nil, err
	}
	sections, err := getAllSectionsAsMap(ctx, repo)
	if err != nil {
		return nil, nil, nil, err
	}
	tasks, err := getAllTasksAsMap(ctx, repo)
	if err != nil {
		return nil, nil, nil, err
	}

	return technologies, sections, tasks, nil
}

func getAllTechnologiesAsMap(
	ctx context.Context,
	repo port.ContentRepository,
) (map[string]model.Technology, error) {
	technologies, err := repo.GetAllTechnologies(ctx)
	if err != nil {
		return nil, err
	}
	technologiesMap := make(map[string]model.Technology)
	for _, technology := range technologies {
		technologiesMap[technology.Id] = technology
	}

	return technologiesMap, nil
}

func getAllSectionsAsMap(
	ctx context.Context,
	repo port.ContentRepository,
) (map[string]model.Section, error) {
	sections, err := repo.GetAllSections(ctx)
	if err != nil {
		return nil, err
	}
	sectionsMap := make(map[string]model.Section)
	for _, section := range sections {
		sectionsMap[section.Id] = section
	}

	return sectionsMap, nil
}

func getAllTasksAsMap(
	ctx context.Context,
	repo port.ContentRepository,
) (map[string]model.Task, error) {
	tasks, err := repo.GetAllTasks(ctx)
	if err != nil {
		return nil, err
	}
	tasksMap := make(map[string]model.Task)
	for _, task := range tasks {
		tasksMap[task.Id] = task
	}

	return tasksMap, nil
}
