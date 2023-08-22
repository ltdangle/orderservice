# Project-X: Order Service
Prjctr Go course. Mock order service project.

## Functionality
- creates order - "/create" - DONE!
- adds/deletes product to the order - "/add-product", "/product" - DONE!
- validates that products exist in inventory (product service)
- calculates order total
- transfers to payment gateway - "/order/{uuid}/checkout" - DONE!
- confirms order payment from payment gateway
- provides order status (as requested from user service, product service, others) - "/retrieve/{uuid}" - DONE!

## Technical implementation
- REST API
- uses repository pattern for entity persistence
- strives to follow clean architecture practices (decoupling via interfaces)

## Deployment
- deployed as binary
- webservice listens on port :8081
