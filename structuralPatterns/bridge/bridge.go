package bridge

// What
//  - is a struc­tur­al design pat­tern
//  - split a large class or a set of close­ly relat­ed class­es
//	  into two sep­a­rate hier­ar­chies (2 hệ phân cấp riêng biệt)
//    —abstrac­tion and imple­men­ta­tion—(phần trừu tượng và phần triển khai)
//    which (để) can be devel­oped inde­pen­dent­ly (độc lập) of each other

// Where
//  - The Bridge pat­tern attempts to solve this prob­lem by switch­ing from
//    inher­i­tance to the object com­po­si­tion
//  - extract one of the dimen­sions into a sep­a­rate class hier­ar­chy
//    so that the orig­i­nal class­es will ref­er­ence an object of the new hier­ar­chy,
//    instead of hav­ing all of its state and behav­iors with­in one class.

// When
//  - avoid com­plex­ity of a large class hier­ar­chy (tránh sự phức tạp của một hệ phân cấp lớp lớn)
//  - want to change and extend the abstraction
//  - want decouple an abstraction from its implementation

// Why
//  - to reduce the amount of subclasses inherited from the base class
//    when adding new functionality (khi thêm chức năng mới)

// How
//  - identify the abstraction and the implementation
//  - create an interface for the abstraction
//  - implement the abstraction interface in a concrete class

type Data struct{}

// Abstraction

type DataParser interface {
	Parse() (*Data, error)
}

type DataPersistent interface {
	Save(data *Data) error
}

func parseAndSaveData(parser DataParser, storage DataPersistent) error {
	data, err := parser.Parse()

	if err != nil {
		return err
	}

	if err := storage.Save(data); err != nil {
		return err
	}

	return nil
}

// Implementation

type MySQLParser struct{}
type MongoParser struct{}
type FileParser struct{}

func (MySQLParser) Parse() (*Data, error) { return &Data{}, nil }
func (MongoParser) Parse() (*Data, error) { return &Data{}, nil }
func (FileParser) Parse() (*Data, error)  { return &Data{}, nil }

type JSONFilePersistent struct{}
type RPCServicePersistent struct{}
type AWSSS3Persistent struct{}

func (JSONFilePersistent) Save(data *Data) error   { return nil }
func (RPCServicePersistent) Save(data *Data) error { return nil }
func (AWSSS3Persistent) Save(data *Data) error     { return nil }

func Caller() {
	_ = parseAndSaveData(MySQLParser{}, JSONFilePersistent{})
	_ = parseAndSaveData(MongoParser{}, RPCServicePersistent{})
	_ = parseAndSaveData(FileParser{}, AWSSS3Persistent{})
}
