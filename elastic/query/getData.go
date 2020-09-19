package query

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

func queryData(ctx context.Context, esclient *elastic.Client, name, category string) {
	if category != "" {
		queryEvent(ctx, esclient, name, category)
	} else {
		queryUser(ctx, esclient, name)
	}
}

func queryEvent(ctx context.Context, esclient *elastic.Client, name, category string) {
	multiQuery := elastic.NewMultiMatchQuery(name, "name").Type("phrase_prefix")

	matchQuery := elastic.NewMatchQuery("category", category)
	query := elastic.NewBoolQuery().Must(multiQuery, matchQuery)

	searchResult, err := esclient.Search().Index("search_data").Query(query).Do(ctx)

	if err != nil {
		panic(err)
	}

	for _, hit := range searchResult.Hits.Hits {
		fmt.Println(string(*hit.Source))
	}
}

func queryUser(ctx context.Context, esclient *elastic.Client, name string) {
	multiQuery := elastic.NewMultiMatchQuery(name, "name", "owner").Type("phrase_prefix")

	searchResult, err := esclient.Search().Index("search_data").Query(multiQuery).Do(ctx)

	if err != nil {
		panic(err)
	}

	for _, hit := range searchResult.Hits.Hits {
		fmt.Println(string(*hit.Source))
	}
}
