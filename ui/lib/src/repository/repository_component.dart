import 'package:angular/angular.dart';
import 'package:angular_router/angular_router.dart';
import 'package:gitpods/routes.dart' as global;
import 'package:gitpods/src/gravatar_component.dart';
import 'package:gitpods/src/loading_component.dart';
import 'package:gitpods/src/repository/repository.dart';
import 'package:gitpods/src/repository/repository_service.dart';
import 'package:gitpods/src/repository/repository_tree.dart';
import 'package:gitpods/src/repository/routes.dart';

@Component(
    selector: 'gitpods-repository',
    templateUrl: 'repository_component.html',
    styleUrls: const [
      'repository_component.css'
    ],
    directives: const [
      coreDirectives,
      routerDirectives,
      LoadingComponent,
      GravatarComponent,
    ],
    providers: const [
      RepositoryService,
      ClassProvider(Routes)
    ],
    exports: [
      Routes,
      RoutePaths,
    ])
class RepositoryComponent implements OnActivate {
  RepositoryComponent(this._repositoryService, this.routes);

  final RepositoryService _repositoryService;
  final Routes routes;

  String ownerName;
  Repository repository;
  List<RepositoryTree> tree;

  @override
  void onActivate(RouterState previous, RouterState current) {
    ownerName = current.parameters['owner'];
    String name = current.parameters['name'];

    this._repositoryService.get(ownerName, name).then((RepositoryPage page) {
      this.repository = page.repository;
    });
  }

  String userProfileUrl() =>
      global.RoutePaths.userProfile.toUrl(parameters: {'username': ownerName});

// TODO: Create method to get repository parameters for the following

  String repositoryUrl() => global.RoutePaths.repository.toUrl(parameters: {
        'owner': ownerName,
        'name': repository.name,
      });

  String filesUrl() => RoutePaths.files.toUrl(parameters: {
        'owner': ownerName,
        'name': repository.name,
      });

  String commitsUrl() => RoutePaths.commits.toUrl(parameters: {
        'owner': ownerName,
        'name': repository.name,
      });

  String settingsUrl() => RoutePaths.settings.toUrl(parameters: {
        'owner': ownerName,
        'name': repository.name,
      });
}

class RepositoryPage {
  RepositoryPage(this.repository);

  Repository repository;
}
