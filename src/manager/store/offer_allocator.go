package store

func (zk *ZkStore) CreateOfferAllocatorItem(item *OfferAllocatorItem) error {
	op := &AtomicOp{
		Op:      OP_ADD,
		Entity:  ENTITY_OFFER_ALLOCATOR_ITEM,
		Param1:  item.OfferID,
		Payload: item,
	}

	return zk.Apply(op)
}

func (zk *ZkStore) DeleteOfferAllocatorItem(offerId string) error {
	op := &AtomicOp{
		Op:     OP_REMOVE,
		Entity: ENTITY_OFFER_ALLOCATOR_ITEM,
		Param1: offerId,
	}

	return zk.Apply(op)
}

func (zk *ZkStore) ListOfferallocatorItems() []*OfferAllocatorItem {
	items := make([]*OfferAllocatorItem, 0)
	for _, item := range zk.OfferAllocator {
		items = append(items, item)
	}

	return items
}
