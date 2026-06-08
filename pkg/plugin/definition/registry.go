package definition

import (
	"context"

	"github.com/beckn-one/beckn-onix/pkg/model"
)

type RegistryLookup interface {
	// looks up Registry entry to obtain public keys to validate signature of the incoming message
	Lookup(ctx context.Context, req *model.Subscription) ([]model.Subscription, error)
}

// RegistryMetadataLookup resolves DeDi registry and node records using DeDi-native path parameters.
// All inputs use the DeDi namespace/registry(/recordName) path convention — these are not generic
// Beckn params and are not interchangeable with the subscriberID/keyID used by RegistryLookup.
type RegistryMetadataLookup interface {
	// LookupRegistry fetches registry-level metadata for a DeDi network registry.
	// Input: namespaceIdentifier (DeDi namespace, e.g. "nfh.global"),
	//        registryName (DeDi registry name, e.g. "retail.network.production").
	// Returns registry metadata including manifest URLs used by ManifestLoader.
	LookupRegistry(ctx context.Context, namespaceIdentifier, registryName string) (*model.RegistryMetadata, error)

	// LookupNode fetches the full subscriber record for a DeDi node by its NodeID.
	// Input: nodeID must be a fully-qualified three-part DeDi path in
	//        namespace/registry/recordName format (e.g. "nfh.global/subscribers.beckn.one/bpp.energy.com").
	// Returns a SubscriberRecord with subscriber identity/endpoint data and any node manifest
	// metadata from the same DeDi response. Meta is empty (not an error) when the participant
	// has not yet published a node manifest.
	LookupNode(ctx context.Context, nodeID string) (*model.SubscriberRecord, error)
}

// RegistryLookupProvider initializes a new registry lookup instance.
type RegistryLookupProvider interface {
	New(context.Context, Cache, map[string]string) (RegistryLookup, func() error, error)
}
