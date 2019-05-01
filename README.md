# address-book
## AddressBook restful Web Service to showcase golang microservices capabilities.

This is a project I created and developed to practise golang restful web services. The end-points of this service facilitates the CRUD operations of an address book, which lets users store addresses of people. MongoDb is used as the data store, which is flexible as far as document structure is concerned, ie. it does not enforce a rigid structure on the data that can be stored in a collection. Also the store part is pretty flexible, just by replacing the repository class, it can access any database. 

In addition to the CRUD operations, it has a csv import export endpoint provided for bulk upload of addresses and also for downloads of the same.

MongoDB is the backend database. The instructions to set up the database will be provided later.


Need to write the test classes to do the integration testing of all the end-points and import export function.

Since I have come this far developing this, I may even use it as an actual address book project I have in mind with a web and mobile interface.

Also planning to make it a part of the marketting platform I am planning to develop.
