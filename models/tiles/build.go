package tiles

const BuildTileCategory TileCategory = "BUILD"

type (
	// Used by Jenkins, Gitlab, Teamcity ... as response structure
	BuildTile struct {
		*Tile

		PreviousStatus TileStatus `json:"previousStatus,omitempty"`
		Author         *Author    `json:"author,omitempty"`

		StartedAt         *int64 `json:"startedAt,omitempty"`
		FinishedAt        *int64 `json:"finishedAt,omitempty"`
		Duration          *int64 `json:"duration,omitempty"`
		EstimatedDuration *int64 `json:"estimatedDuration,omitempty"`
	}

	Author struct {
		Name      string `json:"name,omitempty"`
		AvatarUrl string `json:"avatarUrl,omitempty"`
	}
)

func NewBuildTile(t TileType) *BuildTile {
	return &BuildTile{
		Tile: &Tile{
			Category: BuildTileCategory,
			Type:     t,
		},
	}
}
