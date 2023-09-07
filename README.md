# Project-X: Order Service
Prjctr Go course. Mock order service project.

## Functionality
- creates order - "/create" - DONE!
- adds/deletes product to the order - "/add-product", "/product" - DONE!
- validates that products exist in inventory (product service) - DONE! (product service mock)
- calculates order total - DONE!
- transfers to payment gateway - "/order/{uuid}/checkout" - DONE!
- confirms order payment from payment gateway - "/order/{uuid}/payment/{paymentUuid}" - DONE!
- provides order status (as requested from user service, product service, others) - "/retrieve/{uuid}" - DONE!

## TODO
- implement DI with Wire (function type injection, ftw) 
- investigage using GORM raw url for read model retrieveal (null values?)
- implement html ui with htmx
- use type receivers (read-only) for write model
- implement backend data admin (https://go-admin.com, https://github.com/LyricTian/gin-admin)

## Technical implementation
- REST API
- uses repository pattern for entity persistence
- strives to follow clean architecture practices (decoupling via interfaces)

## Deployment
- deployed as binary
- webservice listens on port :8081
