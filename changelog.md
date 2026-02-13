# Changelog

## [Unreleased]

### Added
### Changed
### Fixed
### Removed
### Deprecated

## [1.2.0] - 2026-02-12

### Added
- Task update service
- Task delete service
- Methods were added to return all tasks as well as to filter by status (todo, done, in-progress)
- Methods were added to update the status of a task

## [1.1.0] - 2026-02-08

### Added
- Task registration service with UUID support
- `generals` package with reusable file operations
- Persistent task storage using JSON format

### Changed
- Integrated `github.com/google/uuid v1.6.0` dependency

## [1.0.0] - 2026-02-05

### Added
- Initial CLI setup
- Task add command implementation