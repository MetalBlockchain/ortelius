// (c) 2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avax

import (
	"context"

	"github.com/MetalBlockchain/metalgo/ids"
	"github.com/MetalBlockchain/ortelius/cfg"
	"github.com/MetalBlockchain/ortelius/models"
	"github.com/MetalBlockchain/ortelius/services/indexes/params"
)

func (r *Reader) ListBlocks(ctx context.Context, params *params.ListBlocksParams) (*models.BlockList, error) {
	dbRunner, err := r.conns.DB().NewSession("list_blocks", cfg.RequestTimeout)
	if err != nil {
		return nil, err
	}

	blocks := []*models.Block{}

	_, err = params.Apply(dbRunner.
		Select("id", "type", "parent_id", "chain_id", "created_at").
		From("pvm_blocks")).
		LoadContext(ctx, &blocks)

	if err != nil {
		return nil, err
	}
	return &models.BlockList{Blocks: blocks}, nil
}

func (r *Reader) GetBlock(ctx context.Context, id ids.ID) (*models.Block, error) {
	list, err := r.ListBlocks(ctx, &params.ListBlocksParams{ListParams: params.ListParams{ID: &id}})
	if err != nil || len(list.Blocks) == 0 {
		return nil, err
	}
	return list.Blocks[0], nil
}
