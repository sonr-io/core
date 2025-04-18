// Code generated by protoc-gen-go-cosmos-orm. DO NOT EDIT.

package didv1

import (
	context "context"
	ormlist "cosmossdk.io/orm/model/ormlist"
	ormtable "cosmossdk.io/orm/model/ormtable"
	ormerrors "cosmossdk.io/orm/types/ormerrors"
)

type AuthenticationTable interface {
	Insert(ctx context.Context, authentication *Authentication) error
	Update(ctx context.Context, authentication *Authentication) error
	Save(ctx context.Context, authentication *Authentication) error
	Delete(ctx context.Context, authentication *Authentication) error
	Has(ctx context.Context, did string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, did string) (*Authentication, error)
	HasByControllerSubject(ctx context.Context, controller string, subject string) (found bool, err error)
	// GetByControllerSubject returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByControllerSubject(ctx context.Context, controller string, subject string) (*Authentication, error)
	List(ctx context.Context, prefixKey AuthenticationIndexKey, opts ...ormlist.Option) (AuthenticationIterator, error)
	ListRange(ctx context.Context, from, to AuthenticationIndexKey, opts ...ormlist.Option) (AuthenticationIterator, error)
	DeleteBy(ctx context.Context, prefixKey AuthenticationIndexKey) error
	DeleteRange(ctx context.Context, from, to AuthenticationIndexKey) error

	doNotImplement()
}

type AuthenticationIterator struct {
	ormtable.Iterator
}

func (i AuthenticationIterator) Value() (*Authentication, error) {
	var authentication Authentication
	err := i.UnmarshalMessage(&authentication)
	return &authentication, err
}

type AuthenticationIndexKey interface {
	id() uint32
	values() []interface{}
	authenticationIndexKey()
}

// primary key starting index..
type AuthenticationPrimaryKey = AuthenticationDidIndexKey

type AuthenticationDidIndexKey struct {
	vs []interface{}
}

func (x AuthenticationDidIndexKey) id() uint32              { return 0 }
func (x AuthenticationDidIndexKey) values() []interface{}   { return x.vs }
func (x AuthenticationDidIndexKey) authenticationIndexKey() {}

func (this AuthenticationDidIndexKey) WithDid(did string) AuthenticationDidIndexKey {
	this.vs = []interface{}{did}
	return this
}

type AuthenticationControllerSubjectIndexKey struct {
	vs []interface{}
}

func (x AuthenticationControllerSubjectIndexKey) id() uint32              { return 1 }
func (x AuthenticationControllerSubjectIndexKey) values() []interface{}   { return x.vs }
func (x AuthenticationControllerSubjectIndexKey) authenticationIndexKey() {}

func (this AuthenticationControllerSubjectIndexKey) WithController(controller string) AuthenticationControllerSubjectIndexKey {
	this.vs = []interface{}{controller}
	return this
}

func (this AuthenticationControllerSubjectIndexKey) WithControllerSubject(controller string, subject string) AuthenticationControllerSubjectIndexKey {
	this.vs = []interface{}{controller, subject}
	return this
}

type authenticationTable struct {
	table ormtable.Table
}

func (this authenticationTable) Insert(ctx context.Context, authentication *Authentication) error {
	return this.table.Insert(ctx, authentication)
}

func (this authenticationTable) Update(ctx context.Context, authentication *Authentication) error {
	return this.table.Update(ctx, authentication)
}

func (this authenticationTable) Save(ctx context.Context, authentication *Authentication) error {
	return this.table.Save(ctx, authentication)
}

func (this authenticationTable) Delete(ctx context.Context, authentication *Authentication) error {
	return this.table.Delete(ctx, authentication)
}

func (this authenticationTable) Has(ctx context.Context, did string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, did)
}

func (this authenticationTable) Get(ctx context.Context, did string) (*Authentication, error) {
	var authentication Authentication
	found, err := this.table.PrimaryKey().Get(ctx, &authentication, did)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &authentication, nil
}

func (this authenticationTable) HasByControllerSubject(ctx context.Context, controller string, subject string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		controller,
		subject,
	)
}

func (this authenticationTable) GetByControllerSubject(ctx context.Context, controller string, subject string) (*Authentication, error) {
	var authentication Authentication
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &authentication,
		controller,
		subject,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &authentication, nil
}

func (this authenticationTable) List(ctx context.Context, prefixKey AuthenticationIndexKey, opts ...ormlist.Option) (AuthenticationIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return AuthenticationIterator{it}, err
}

func (this authenticationTable) ListRange(ctx context.Context, from, to AuthenticationIndexKey, opts ...ormlist.Option) (AuthenticationIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return AuthenticationIterator{it}, err
}

func (this authenticationTable) DeleteBy(ctx context.Context, prefixKey AuthenticationIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this authenticationTable) DeleteRange(ctx context.Context, from, to AuthenticationIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this authenticationTable) doNotImplement() {}

var _ AuthenticationTable = authenticationTable{}

func NewAuthenticationTable(db ormtable.Schema) (AuthenticationTable, error) {
	table := db.GetTable(&Authentication{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Authentication{}).ProtoReflect().Descriptor().FullName()))
	}
	return authenticationTable{table}, nil
}

type AssertionTable interface {
	Insert(ctx context.Context, assertion *Assertion) error
	Update(ctx context.Context, assertion *Assertion) error
	Save(ctx context.Context, assertion *Assertion) error
	Delete(ctx context.Context, assertion *Assertion) error
	Has(ctx context.Context, did string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, did string) (*Assertion, error)
	HasByControllerSubject(ctx context.Context, controller string, subject string) (found bool, err error)
	// GetByControllerSubject returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByControllerSubject(ctx context.Context, controller string, subject string) (*Assertion, error)
	List(ctx context.Context, prefixKey AssertionIndexKey, opts ...ormlist.Option) (AssertionIterator, error)
	ListRange(ctx context.Context, from, to AssertionIndexKey, opts ...ormlist.Option) (AssertionIterator, error)
	DeleteBy(ctx context.Context, prefixKey AssertionIndexKey) error
	DeleteRange(ctx context.Context, from, to AssertionIndexKey) error

	doNotImplement()
}

type AssertionIterator struct {
	ormtable.Iterator
}

func (i AssertionIterator) Value() (*Assertion, error) {
	var assertion Assertion
	err := i.UnmarshalMessage(&assertion)
	return &assertion, err
}

type AssertionIndexKey interface {
	id() uint32
	values() []interface{}
	assertionIndexKey()
}

// primary key starting index..
type AssertionPrimaryKey = AssertionDidIndexKey

type AssertionDidIndexKey struct {
	vs []interface{}
}

func (x AssertionDidIndexKey) id() uint32            { return 0 }
func (x AssertionDidIndexKey) values() []interface{} { return x.vs }
func (x AssertionDidIndexKey) assertionIndexKey()    {}

func (this AssertionDidIndexKey) WithDid(did string) AssertionDidIndexKey {
	this.vs = []interface{}{did}
	return this
}

type AssertionControllerSubjectIndexKey struct {
	vs []interface{}
}

func (x AssertionControllerSubjectIndexKey) id() uint32            { return 1 }
func (x AssertionControllerSubjectIndexKey) values() []interface{} { return x.vs }
func (x AssertionControllerSubjectIndexKey) assertionIndexKey()    {}

func (this AssertionControllerSubjectIndexKey) WithController(controller string) AssertionControllerSubjectIndexKey {
	this.vs = []interface{}{controller}
	return this
}

func (this AssertionControllerSubjectIndexKey) WithControllerSubject(controller string, subject string) AssertionControllerSubjectIndexKey {
	this.vs = []interface{}{controller, subject}
	return this
}

type assertionTable struct {
	table ormtable.Table
}

func (this assertionTable) Insert(ctx context.Context, assertion *Assertion) error {
	return this.table.Insert(ctx, assertion)
}

func (this assertionTable) Update(ctx context.Context, assertion *Assertion) error {
	return this.table.Update(ctx, assertion)
}

func (this assertionTable) Save(ctx context.Context, assertion *Assertion) error {
	return this.table.Save(ctx, assertion)
}

func (this assertionTable) Delete(ctx context.Context, assertion *Assertion) error {
	return this.table.Delete(ctx, assertion)
}

func (this assertionTable) Has(ctx context.Context, did string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, did)
}

func (this assertionTable) Get(ctx context.Context, did string) (*Assertion, error) {
	var assertion Assertion
	found, err := this.table.PrimaryKey().Get(ctx, &assertion, did)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &assertion, nil
}

func (this assertionTable) HasByControllerSubject(ctx context.Context, controller string, subject string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		controller,
		subject,
	)
}

func (this assertionTable) GetByControllerSubject(ctx context.Context, controller string, subject string) (*Assertion, error) {
	var assertion Assertion
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &assertion,
		controller,
		subject,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &assertion, nil
}

func (this assertionTable) List(ctx context.Context, prefixKey AssertionIndexKey, opts ...ormlist.Option) (AssertionIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return AssertionIterator{it}, err
}

func (this assertionTable) ListRange(ctx context.Context, from, to AssertionIndexKey, opts ...ormlist.Option) (AssertionIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return AssertionIterator{it}, err
}

func (this assertionTable) DeleteBy(ctx context.Context, prefixKey AssertionIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this assertionTable) DeleteRange(ctx context.Context, from, to AssertionIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this assertionTable) doNotImplement() {}

var _ AssertionTable = assertionTable{}

func NewAssertionTable(db ormtable.Schema) (AssertionTable, error) {
	table := db.GetTable(&Assertion{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Assertion{}).ProtoReflect().Descriptor().FullName()))
	}
	return assertionTable{table}, nil
}

type ControllerTable interface {
	Insert(ctx context.Context, controller *Controller) error
	Update(ctx context.Context, controller *Controller) error
	Save(ctx context.Context, controller *Controller) error
	Delete(ctx context.Context, controller *Controller) error
	Has(ctx context.Context, did string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, did string) (*Controller, error)
	HasByAddress(ctx context.Context, address string) (found bool, err error)
	// GetByAddress returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByAddress(ctx context.Context, address string) (*Controller, error)
	HasBySubject(ctx context.Context, subject string) (found bool, err error)
	// GetBySubject returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetBySubject(ctx context.Context, subject string) (*Controller, error)
	HasByPublicKeyBase64(ctx context.Context, public_key_base64 string) (found bool, err error)
	// GetByPublicKeyBase64 returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByPublicKeyBase64(ctx context.Context, public_key_base64 string) (*Controller, error)
	List(ctx context.Context, prefixKey ControllerIndexKey, opts ...ormlist.Option) (ControllerIterator, error)
	ListRange(ctx context.Context, from, to ControllerIndexKey, opts ...ormlist.Option) (ControllerIterator, error)
	DeleteBy(ctx context.Context, prefixKey ControllerIndexKey) error
	DeleteRange(ctx context.Context, from, to ControllerIndexKey) error

	doNotImplement()
}

type ControllerIterator struct {
	ormtable.Iterator
}

func (i ControllerIterator) Value() (*Controller, error) {
	var controller Controller
	err := i.UnmarshalMessage(&controller)
	return &controller, err
}

type ControllerIndexKey interface {
	id() uint32
	values() []interface{}
	controllerIndexKey()
}

// primary key starting index..
type ControllerPrimaryKey = ControllerDidIndexKey

type ControllerDidIndexKey struct {
	vs []interface{}
}

func (x ControllerDidIndexKey) id() uint32            { return 0 }
func (x ControllerDidIndexKey) values() []interface{} { return x.vs }
func (x ControllerDidIndexKey) controllerIndexKey()   {}

func (this ControllerDidIndexKey) WithDid(did string) ControllerDidIndexKey {
	this.vs = []interface{}{did}
	return this
}

type ControllerAddressIndexKey struct {
	vs []interface{}
}

func (x ControllerAddressIndexKey) id() uint32            { return 1 }
func (x ControllerAddressIndexKey) values() []interface{} { return x.vs }
func (x ControllerAddressIndexKey) controllerIndexKey()   {}

func (this ControllerAddressIndexKey) WithAddress(address string) ControllerAddressIndexKey {
	this.vs = []interface{}{address}
	return this
}

type ControllerSubjectIndexKey struct {
	vs []interface{}
}

func (x ControllerSubjectIndexKey) id() uint32            { return 2 }
func (x ControllerSubjectIndexKey) values() []interface{} { return x.vs }
func (x ControllerSubjectIndexKey) controllerIndexKey()   {}

func (this ControllerSubjectIndexKey) WithSubject(subject string) ControllerSubjectIndexKey {
	this.vs = []interface{}{subject}
	return this
}

type ControllerPublicKeyBase64IndexKey struct {
	vs []interface{}
}

func (x ControllerPublicKeyBase64IndexKey) id() uint32            { return 3 }
func (x ControllerPublicKeyBase64IndexKey) values() []interface{} { return x.vs }
func (x ControllerPublicKeyBase64IndexKey) controllerIndexKey()   {}

func (this ControllerPublicKeyBase64IndexKey) WithPublicKeyBase64(public_key_base64 string) ControllerPublicKeyBase64IndexKey {
	this.vs = []interface{}{public_key_base64}
	return this
}

type controllerTable struct {
	table ormtable.Table
}

func (this controllerTable) Insert(ctx context.Context, controller *Controller) error {
	return this.table.Insert(ctx, controller)
}

func (this controllerTable) Update(ctx context.Context, controller *Controller) error {
	return this.table.Update(ctx, controller)
}

func (this controllerTable) Save(ctx context.Context, controller *Controller) error {
	return this.table.Save(ctx, controller)
}

func (this controllerTable) Delete(ctx context.Context, controller *Controller) error {
	return this.table.Delete(ctx, controller)
}

func (this controllerTable) Has(ctx context.Context, did string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, did)
}

func (this controllerTable) Get(ctx context.Context, did string) (*Controller, error) {
	var controller Controller
	found, err := this.table.PrimaryKey().Get(ctx, &controller, did)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &controller, nil
}

func (this controllerTable) HasByAddress(ctx context.Context, address string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		address,
	)
}

func (this controllerTable) GetByAddress(ctx context.Context, address string) (*Controller, error) {
	var controller Controller
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &controller,
		address,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &controller, nil
}

func (this controllerTable) HasBySubject(ctx context.Context, subject string) (found bool, err error) {
	return this.table.GetIndexByID(2).(ormtable.UniqueIndex).Has(ctx,
		subject,
	)
}

func (this controllerTable) GetBySubject(ctx context.Context, subject string) (*Controller, error) {
	var controller Controller
	found, err := this.table.GetIndexByID(2).(ormtable.UniqueIndex).Get(ctx, &controller,
		subject,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &controller, nil
}

func (this controllerTable) HasByPublicKeyBase64(ctx context.Context, public_key_base64 string) (found bool, err error) {
	return this.table.GetIndexByID(3).(ormtable.UniqueIndex).Has(ctx,
		public_key_base64,
	)
}

func (this controllerTable) GetByPublicKeyBase64(ctx context.Context, public_key_base64 string) (*Controller, error) {
	var controller Controller
	found, err := this.table.GetIndexByID(3).(ormtable.UniqueIndex).Get(ctx, &controller,
		public_key_base64,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &controller, nil
}

func (this controllerTable) List(ctx context.Context, prefixKey ControllerIndexKey, opts ...ormlist.Option) (ControllerIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return ControllerIterator{it}, err
}

func (this controllerTable) ListRange(ctx context.Context, from, to ControllerIndexKey, opts ...ormlist.Option) (ControllerIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return ControllerIterator{it}, err
}

func (this controllerTable) DeleteBy(ctx context.Context, prefixKey ControllerIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this controllerTable) DeleteRange(ctx context.Context, from, to ControllerIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this controllerTable) doNotImplement() {}

var _ ControllerTable = controllerTable{}

func NewControllerTable(db ormtable.Schema) (ControllerTable, error) {
	table := db.GetTable(&Controller{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Controller{}).ProtoReflect().Descriptor().FullName()))
	}
	return controllerTable{table}, nil
}

type DelegationTable interface {
	Insert(ctx context.Context, delegation *Delegation) error
	Update(ctx context.Context, delegation *Delegation) error
	Save(ctx context.Context, delegation *Delegation) error
	Delete(ctx context.Context, delegation *Delegation) error
	Has(ctx context.Context, did string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, did string) (*Delegation, error)
	HasByControllerSubject(ctx context.Context, controller string, subject string) (found bool, err error)
	// GetByControllerSubject returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByControllerSubject(ctx context.Context, controller string, subject string) (*Delegation, error)
	List(ctx context.Context, prefixKey DelegationIndexKey, opts ...ormlist.Option) (DelegationIterator, error)
	ListRange(ctx context.Context, from, to DelegationIndexKey, opts ...ormlist.Option) (DelegationIterator, error)
	DeleteBy(ctx context.Context, prefixKey DelegationIndexKey) error
	DeleteRange(ctx context.Context, from, to DelegationIndexKey) error

	doNotImplement()
}

type DelegationIterator struct {
	ormtable.Iterator
}

func (i DelegationIterator) Value() (*Delegation, error) {
	var delegation Delegation
	err := i.UnmarshalMessage(&delegation)
	return &delegation, err
}

type DelegationIndexKey interface {
	id() uint32
	values() []interface{}
	delegationIndexKey()
}

// primary key starting index..
type DelegationPrimaryKey = DelegationDidIndexKey

type DelegationDidIndexKey struct {
	vs []interface{}
}

func (x DelegationDidIndexKey) id() uint32            { return 0 }
func (x DelegationDidIndexKey) values() []interface{} { return x.vs }
func (x DelegationDidIndexKey) delegationIndexKey()   {}

func (this DelegationDidIndexKey) WithDid(did string) DelegationDidIndexKey {
	this.vs = []interface{}{did}
	return this
}

type DelegationControllerSubjectIndexKey struct {
	vs []interface{}
}

func (x DelegationControllerSubjectIndexKey) id() uint32            { return 1 }
func (x DelegationControllerSubjectIndexKey) values() []interface{} { return x.vs }
func (x DelegationControllerSubjectIndexKey) delegationIndexKey()   {}

func (this DelegationControllerSubjectIndexKey) WithController(controller string) DelegationControllerSubjectIndexKey {
	this.vs = []interface{}{controller}
	return this
}

func (this DelegationControllerSubjectIndexKey) WithControllerSubject(controller string, subject string) DelegationControllerSubjectIndexKey {
	this.vs = []interface{}{controller, subject}
	return this
}

type delegationTable struct {
	table ormtable.Table
}

func (this delegationTable) Insert(ctx context.Context, delegation *Delegation) error {
	return this.table.Insert(ctx, delegation)
}

func (this delegationTable) Update(ctx context.Context, delegation *Delegation) error {
	return this.table.Update(ctx, delegation)
}

func (this delegationTable) Save(ctx context.Context, delegation *Delegation) error {
	return this.table.Save(ctx, delegation)
}

func (this delegationTable) Delete(ctx context.Context, delegation *Delegation) error {
	return this.table.Delete(ctx, delegation)
}

func (this delegationTable) Has(ctx context.Context, did string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, did)
}

func (this delegationTable) Get(ctx context.Context, did string) (*Delegation, error) {
	var delegation Delegation
	found, err := this.table.PrimaryKey().Get(ctx, &delegation, did)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &delegation, nil
}

func (this delegationTable) HasByControllerSubject(ctx context.Context, controller string, subject string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		controller,
		subject,
	)
}

func (this delegationTable) GetByControllerSubject(ctx context.Context, controller string, subject string) (*Delegation, error) {
	var delegation Delegation
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &delegation,
		controller,
		subject,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &delegation, nil
}

func (this delegationTable) List(ctx context.Context, prefixKey DelegationIndexKey, opts ...ormlist.Option) (DelegationIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return DelegationIterator{it}, err
}

func (this delegationTable) ListRange(ctx context.Context, from, to DelegationIndexKey, opts ...ormlist.Option) (DelegationIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return DelegationIterator{it}, err
}

func (this delegationTable) DeleteBy(ctx context.Context, prefixKey DelegationIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this delegationTable) DeleteRange(ctx context.Context, from, to DelegationIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this delegationTable) doNotImplement() {}

var _ DelegationTable = delegationTable{}

func NewDelegationTable(db ormtable.Schema) (DelegationTable, error) {
	table := db.GetTable(&Delegation{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Delegation{}).ProtoReflect().Descriptor().FullName()))
	}
	return delegationTable{table}, nil
}

type InvocationTable interface {
	Insert(ctx context.Context, invocation *Invocation) error
	Update(ctx context.Context, invocation *Invocation) error
	Save(ctx context.Context, invocation *Invocation) error
	Delete(ctx context.Context, invocation *Invocation) error
	Has(ctx context.Context, did string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, did string) (*Invocation, error)
	HasByControllerSubject(ctx context.Context, controller string, subject string) (found bool, err error)
	// GetByControllerSubject returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByControllerSubject(ctx context.Context, controller string, subject string) (*Invocation, error)
	List(ctx context.Context, prefixKey InvocationIndexKey, opts ...ormlist.Option) (InvocationIterator, error)
	ListRange(ctx context.Context, from, to InvocationIndexKey, opts ...ormlist.Option) (InvocationIterator, error)
	DeleteBy(ctx context.Context, prefixKey InvocationIndexKey) error
	DeleteRange(ctx context.Context, from, to InvocationIndexKey) error

	doNotImplement()
}

type InvocationIterator struct {
	ormtable.Iterator
}

func (i InvocationIterator) Value() (*Invocation, error) {
	var invocation Invocation
	err := i.UnmarshalMessage(&invocation)
	return &invocation, err
}

type InvocationIndexKey interface {
	id() uint32
	values() []interface{}
	invocationIndexKey()
}

// primary key starting index..
type InvocationPrimaryKey = InvocationDidIndexKey

type InvocationDidIndexKey struct {
	vs []interface{}
}

func (x InvocationDidIndexKey) id() uint32            { return 0 }
func (x InvocationDidIndexKey) values() []interface{} { return x.vs }
func (x InvocationDidIndexKey) invocationIndexKey()   {}

func (this InvocationDidIndexKey) WithDid(did string) InvocationDidIndexKey {
	this.vs = []interface{}{did}
	return this
}

type InvocationControllerSubjectIndexKey struct {
	vs []interface{}
}

func (x InvocationControllerSubjectIndexKey) id() uint32            { return 1 }
func (x InvocationControllerSubjectIndexKey) values() []interface{} { return x.vs }
func (x InvocationControllerSubjectIndexKey) invocationIndexKey()   {}

func (this InvocationControllerSubjectIndexKey) WithController(controller string) InvocationControllerSubjectIndexKey {
	this.vs = []interface{}{controller}
	return this
}

func (this InvocationControllerSubjectIndexKey) WithControllerSubject(controller string, subject string) InvocationControllerSubjectIndexKey {
	this.vs = []interface{}{controller, subject}
	return this
}

type invocationTable struct {
	table ormtable.Table
}

func (this invocationTable) Insert(ctx context.Context, invocation *Invocation) error {
	return this.table.Insert(ctx, invocation)
}

func (this invocationTable) Update(ctx context.Context, invocation *Invocation) error {
	return this.table.Update(ctx, invocation)
}

func (this invocationTable) Save(ctx context.Context, invocation *Invocation) error {
	return this.table.Save(ctx, invocation)
}

func (this invocationTable) Delete(ctx context.Context, invocation *Invocation) error {
	return this.table.Delete(ctx, invocation)
}

func (this invocationTable) Has(ctx context.Context, did string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, did)
}

func (this invocationTable) Get(ctx context.Context, did string) (*Invocation, error) {
	var invocation Invocation
	found, err := this.table.PrimaryKey().Get(ctx, &invocation, did)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &invocation, nil
}

func (this invocationTable) HasByControllerSubject(ctx context.Context, controller string, subject string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		controller,
		subject,
	)
}

func (this invocationTable) GetByControllerSubject(ctx context.Context, controller string, subject string) (*Invocation, error) {
	var invocation Invocation
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &invocation,
		controller,
		subject,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &invocation, nil
}

func (this invocationTable) List(ctx context.Context, prefixKey InvocationIndexKey, opts ...ormlist.Option) (InvocationIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return InvocationIterator{it}, err
}

func (this invocationTable) ListRange(ctx context.Context, from, to InvocationIndexKey, opts ...ormlist.Option) (InvocationIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return InvocationIterator{it}, err
}

func (this invocationTable) DeleteBy(ctx context.Context, prefixKey InvocationIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this invocationTable) DeleteRange(ctx context.Context, from, to InvocationIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this invocationTable) doNotImplement() {}

var _ InvocationTable = invocationTable{}

func NewInvocationTable(db ormtable.Schema) (InvocationTable, error) {
	table := db.GetTable(&Invocation{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Invocation{}).ProtoReflect().Descriptor().FullName()))
	}
	return invocationTable{table}, nil
}

type StateStore interface {
	AuthenticationTable() AuthenticationTable
	AssertionTable() AssertionTable
	ControllerTable() ControllerTable
	DelegationTable() DelegationTable
	InvocationTable() InvocationTable

	doNotImplement()
}

type stateStore struct {
	authentication AuthenticationTable
	assertion      AssertionTable
	controller     ControllerTable
	delegation     DelegationTable
	invocation     InvocationTable
}

func (x stateStore) AuthenticationTable() AuthenticationTable {
	return x.authentication
}

func (x stateStore) AssertionTable() AssertionTable {
	return x.assertion
}

func (x stateStore) ControllerTable() ControllerTable {
	return x.controller
}

func (x stateStore) DelegationTable() DelegationTable {
	return x.delegation
}

func (x stateStore) InvocationTable() InvocationTable {
	return x.invocation
}

func (stateStore) doNotImplement() {}

var _ StateStore = stateStore{}

func NewStateStore(db ormtable.Schema) (StateStore, error) {
	authenticationTable, err := NewAuthenticationTable(db)
	if err != nil {
		return nil, err
	}

	assertionTable, err := NewAssertionTable(db)
	if err != nil {
		return nil, err
	}

	controllerTable, err := NewControllerTable(db)
	if err != nil {
		return nil, err
	}

	delegationTable, err := NewDelegationTable(db)
	if err != nil {
		return nil, err
	}

	invocationTable, err := NewInvocationTable(db)
	if err != nil {
		return nil, err
	}

	return stateStore{
		authenticationTable,
		assertionTable,
		controllerTable,
		delegationTable,
		invocationTable,
	}, nil
}
