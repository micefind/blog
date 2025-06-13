import preventClick from './preventClick'

export default function setupDirectives(app) {
  app.directive('prevent-click', preventClick)
}
