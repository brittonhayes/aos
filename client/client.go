// Code generated by github.com/Yamashou/gqlgenc, DO NOT EDIT.

package client

import (
	"context"
	"net/http"

	"github.com/Yamashou/gqlgenc/clientv2"
	"github.com/brittonhayes/aos/api"
)

type Client struct {
	Client *clientv2.Client
}

func NewClient(cli *http.Client, baseURL string, options *clientv2.Options, interceptors ...clientv2.RequestInterceptor) *Client {
	return &Client{Client: clientv2.NewClient(cli, baseURL, options, interceptors...)}
}

type Query struct {
	Allegiances    []*api.Allegiance    "json:\"allegiances\" graphql:\"allegiances\""
	GrandAlliances []*api.GrandAlliance "json:\"grandAlliances\" graphql:\"grandAlliances\""
	Units          []*api.Unit          "json:\"units\" graphql:\"units\""
	Warscrolls     []*api.Warscroll     "json:\"warscrolls\" graphql:\"warscrolls\""
}
type UnitFragment struct {
	ID             string                         "json:\"id\" graphql:\"id\""
	Name           string                         "json:\"name\" graphql:\"name\""
	GrandAlliance  *string                        "json:\"grandAlliance,omitempty\" graphql:\"grandAlliance\""
	Champion       *string                        "json:\"champion,omitempty\" graphql:\"champion\""
	Size           *string                        "json:\"size,omitempty\" graphql:\"size\""
	Move           *string                        "json:\"move,omitempty\" graphql:\"move\""
	Description    *string                        "json:\"description,omitempty\" graphql:\"description\""
	Save           *int64                         "json:\"save,omitempty\" graphql:\"save\""
	Bravery        *int64                         "json:\"bravery,omitempty\" graphql:\"bravery\""
	Models         *int64                         "json:\"models,omitempty\" graphql:\"models\""
	Points         *int64                         "json:\"points,omitempty\" graphql:\"points\""
	Wounds         *int64                         "json:\"wounds,omitempty\" graphql:\"wounds\""
	Keywords       []*string                      "json:\"keywords,omitempty\" graphql:\"keywords\""
	MissileWeapons []*UnitFragment_MissileWeapons "json:\"missileWeapons,omitempty\" graphql:\"missileWeapons\""
	MeleeWeapons   []*UnitFragment_MeleeWeapons   "json:\"meleeWeapons,omitempty\" graphql:\"meleeWeapons\""
}

func (t *UnitFragment) GetID() string {
	if t == nil {
		t = &UnitFragment{}
	}
	return t.ID
}
func (t *UnitFragment) GetName() string {
	if t == nil {
		t = &UnitFragment{}
	}
	return t.Name
}
func (t *UnitFragment) GetGrandAlliance() *string {
	if t == nil {
		t = &UnitFragment{}
	}
	return t.GrandAlliance
}
func (t *UnitFragment) GetChampion() *string {
	if t == nil {
		t = &UnitFragment{}
	}
	return t.Champion
}
func (t *UnitFragment) GetSize() *string {
	if t == nil {
		t = &UnitFragment{}
	}
	return t.Size
}
func (t *UnitFragment) GetMove() *string {
	if t == nil {
		t = &UnitFragment{}
	}
	return t.Move
}
func (t *UnitFragment) GetDescription() *string {
	if t == nil {
		t = &UnitFragment{}
	}
	return t.Description
}
func (t *UnitFragment) GetSave() *int64 {
	if t == nil {
		t = &UnitFragment{}
	}
	return t.Save
}
func (t *UnitFragment) GetBravery() *int64 {
	if t == nil {
		t = &UnitFragment{}
	}
	return t.Bravery
}
func (t *UnitFragment) GetModels() *int64 {
	if t == nil {
		t = &UnitFragment{}
	}
	return t.Models
}
func (t *UnitFragment) GetPoints() *int64 {
	if t == nil {
		t = &UnitFragment{}
	}
	return t.Points
}
func (t *UnitFragment) GetWounds() *int64 {
	if t == nil {
		t = &UnitFragment{}
	}
	return t.Wounds
}
func (t *UnitFragment) GetKeywords() []*string {
	if t == nil {
		t = &UnitFragment{}
	}
	return t.Keywords
}
func (t *UnitFragment) GetMissileWeapons() []*UnitFragment_MissileWeapons {
	if t == nil {
		t = &UnitFragment{}
	}
	return t.MissileWeapons
}
func (t *UnitFragment) GetMeleeWeapons() []*UnitFragment_MeleeWeapons {
	if t == nil {
		t = &UnitFragment{}
	}
	return t.MeleeWeapons
}

type AllegianceFragment struct {
	ID            string  "json:\"id\" graphql:\"id\""
	Name          string  "json:\"name\" graphql:\"name\""
	Description   *string "json:\"description,omitempty\" graphql:\"description\""
	GrandAlliance *string "json:\"grandAlliance,omitempty\" graphql:\"grandAlliance\""
	MortalRealm   *string "json:\"mortalRealm,omitempty\" graphql:\"mortalRealm\""
}

func (t *AllegianceFragment) GetID() string {
	if t == nil {
		t = &AllegianceFragment{}
	}
	return t.ID
}
func (t *AllegianceFragment) GetName() string {
	if t == nil {
		t = &AllegianceFragment{}
	}
	return t.Name
}
func (t *AllegianceFragment) GetDescription() *string {
	if t == nil {
		t = &AllegianceFragment{}
	}
	return t.Description
}
func (t *AllegianceFragment) GetGrandAlliance() *string {
	if t == nil {
		t = &AllegianceFragment{}
	}
	return t.GrandAlliance
}
func (t *AllegianceFragment) GetMortalRealm() *string {
	if t == nil {
		t = &AllegianceFragment{}
	}
	return t.MortalRealm
}

type GrandAllianceFragment struct {
	ID          string  "json:\"id\" graphql:\"id\""
	Name        string  "json:\"name\" graphql:\"name\""
	Description *string "json:\"description,omitempty\" graphql:\"description\""
}

func (t *GrandAllianceFragment) GetID() string {
	if t == nil {
		t = &GrandAllianceFragment{}
	}
	return t.ID
}
func (t *GrandAllianceFragment) GetName() string {
	if t == nil {
		t = &GrandAllianceFragment{}
	}
	return t.Name
}
func (t *GrandAllianceFragment) GetDescription() *string {
	if t == nil {
		t = &GrandAllianceFragment{}
	}
	return t.Description
}

type WarscrollFragment struct {
	ID              string  "json:\"id\" graphql:\"id\""
	Name            string  "json:\"name\" graphql:\"name\""
	AllegianceID    *string "json:\"allegianceId,omitempty\" graphql:\"allegianceId\""
	GrandAllianceID *string "json:\"grandAllianceId,omitempty\" graphql:\"grandAllianceId\""
	Size            *int64  "json:\"size,omitempty\" graphql:\"size\""
	Points          *int64  "json:\"points,omitempty\" graphql:\"points\""
	BattlefieldRole *string "json:\"battlefieldRole,omitempty\" graphql:\"battlefieldRole\""
	Notes           *string "json:\"notes,omitempty\" graphql:\"notes\""
}

func (t *WarscrollFragment) GetID() string {
	if t == nil {
		t = &WarscrollFragment{}
	}
	return t.ID
}
func (t *WarscrollFragment) GetName() string {
	if t == nil {
		t = &WarscrollFragment{}
	}
	return t.Name
}
func (t *WarscrollFragment) GetAllegianceID() *string {
	if t == nil {
		t = &WarscrollFragment{}
	}
	return t.AllegianceID
}
func (t *WarscrollFragment) GetGrandAllianceID() *string {
	if t == nil {
		t = &WarscrollFragment{}
	}
	return t.GrandAllianceID
}
func (t *WarscrollFragment) GetSize() *int64 {
	if t == nil {
		t = &WarscrollFragment{}
	}
	return t.Size
}
func (t *WarscrollFragment) GetPoints() *int64 {
	if t == nil {
		t = &WarscrollFragment{}
	}
	return t.Points
}
func (t *WarscrollFragment) GetBattlefieldRole() *string {
	if t == nil {
		t = &WarscrollFragment{}
	}
	return t.BattlefieldRole
}
func (t *WarscrollFragment) GetNotes() *string {
	if t == nil {
		t = &WarscrollFragment{}
	}
	return t.Notes
}

type UnitFragment_MissileWeapons struct {
	ID string "json:\"id\" graphql:\"id\""
}

func (t *UnitFragment_MissileWeapons) GetID() string {
	if t == nil {
		t = &UnitFragment_MissileWeapons{}
	}
	return t.ID
}

type UnitFragment_MeleeWeapons struct {
	ID string "json:\"id\" graphql:\"id\""
}

func (t *UnitFragment_MeleeWeapons) GetID() string {
	if t == nil {
		t = &UnitFragment_MeleeWeapons{}
	}
	return t.ID
}

type GetUnits_Units_UnitFragment_MissileWeapons struct {
	ID string "json:\"id\" graphql:\"id\""
}

func (t *GetUnits_Units_UnitFragment_MissileWeapons) GetID() string {
	if t == nil {
		t = &GetUnits_Units_UnitFragment_MissileWeapons{}
	}
	return t.ID
}

type GetUnits_Units_UnitFragment_MeleeWeapons struct {
	ID string "json:\"id\" graphql:\"id\""
}

func (t *GetUnits_Units_UnitFragment_MeleeWeapons) GetID() string {
	if t == nil {
		t = &GetUnits_Units_UnitFragment_MeleeWeapons{}
	}
	return t.ID
}

type GetUnits struct {
	Units []*UnitFragment "json:\"units\" graphql:\"units\""
}

func (t *GetUnits) GetUnits() []*UnitFragment {
	if t == nil {
		t = &GetUnits{}
	}
	return t.Units
}

type GetAllegiances struct {
	Allegiances []*AllegianceFragment "json:\"allegiances\" graphql:\"allegiances\""
}

func (t *GetAllegiances) GetAllegiances() []*AllegianceFragment {
	if t == nil {
		t = &GetAllegiances{}
	}
	return t.Allegiances
}

type GetGrandAlliances struct {
	GrandAlliances []*GrandAllianceFragment "json:\"grandAlliances\" graphql:\"grandAlliances\""
}

func (t *GetGrandAlliances) GetGrandAlliances() []*GrandAllianceFragment {
	if t == nil {
		t = &GetGrandAlliances{}
	}
	return t.GrandAlliances
}

type GetWarscrolls struct {
	Warscrolls []*WarscrollFragment "json:\"warscrolls\" graphql:\"warscrolls\""
}

func (t *GetWarscrolls) GetWarscrolls() []*WarscrollFragment {
	if t == nil {
		t = &GetWarscrolls{}
	}
	return t.Warscrolls
}

const GetUnitsDocument = `query GetUnits ($filter: UnitFilters!) {
	units(filter: $filter) {
		... UnitFragment
	}
}
fragment UnitFragment on Unit {
	id
	name
	grandAlliance
	champion
	size
	move
	description
	save
	bravery
	models
	points
	wounds
	keywords
	missileWeapons {
		id
	}
	meleeWeapons {
		id
	}
}
`

func (c *Client) GetUnits(ctx context.Context, filter UnitFilters, interceptors ...clientv2.RequestInterceptor) (*GetUnits, error) {
	vars := map[string]interface{}{
		"filter": filter,
	}

	var res GetUnits
	if err := c.Client.Post(ctx, "GetUnits", GetUnitsDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetAllegiancesDocument = `query GetAllegiances ($filter: AllegianceFilters!) {
	allegiances(filter: $filter) {
		... AllegianceFragment
	}
}
fragment AllegianceFragment on Allegiance {
	id
	name
	description
	grandAlliance
	mortalRealm
}
`

func (c *Client) GetAllegiances(ctx context.Context, filter AllegianceFilters, interceptors ...clientv2.RequestInterceptor) (*GetAllegiances, error) {
	vars := map[string]interface{}{
		"filter": filter,
	}

	var res GetAllegiances
	if err := c.Client.Post(ctx, "GetAllegiances", GetAllegiancesDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetGrandAlliancesDocument = `query GetGrandAlliances {
	grandAlliances {
		... GrandAllianceFragment
	}
}
fragment GrandAllianceFragment on GrandAlliance {
	id
	name
	description
}
`

func (c *Client) GetGrandAlliances(ctx context.Context, interceptors ...clientv2.RequestInterceptor) (*GetGrandAlliances, error) {
	vars := map[string]interface{}{}

	var res GetGrandAlliances
	if err := c.Client.Post(ctx, "GetGrandAlliances", GetGrandAlliancesDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetWarscrollsDocument = `query GetWarscrolls ($filter: WarscrollFilters!) {
	warscrolls(filter: $filter) {
		... WarscrollFragment
	}
}
fragment WarscrollFragment on Warscroll {
	id
	name
	allegianceId
	grandAllianceId
	size
	points
	battlefieldRole
	notes
}
`

func (c *Client) GetWarscrolls(ctx context.Context, filter WarscrollFilters, interceptors ...clientv2.RequestInterceptor) (*GetWarscrolls, error) {
	vars := map[string]interface{}{
		"filter": filter,
	}

	var res GetWarscrolls
	if err := c.Client.Post(ctx, "GetWarscrolls", GetWarscrollsDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}