# Project-X: Order Service
Prjctr Go course. Mock order service project.

## Functionality
- creates order
- adds/deletes product to the order
- validates that products exist in inventory (product service)
- calculates order total
- transfers to payment gateway
- confirms order payment from payment gateway
- provides order status (as requested from user service, product service, others)

## Technical implementation
- REST API
- uses repository pattern for entity persistence
- strives to follow clean architecture practices (decoupling via interfaces)
