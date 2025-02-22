%global project	https://github.com/kirici/myapp

Name:           myapp
Version:        1.6.1
Release:        1%{?dist}
Summary:        Placeholder application

License:        GPLv3
URL:            %{project}
Source0:        %{project}/archive/refs/tags/v%{version}.tar.gz

BuildRequires:  golang
BuildRequires:  git

Provides:       %{name} = %{version}

%description
Placeholder application - minimal webserver in Go using Gin and exporting Prometheus metrics.

%global debug_package %{nil}

%prep
%autosetup


%build
go build -v -o %{name}


%install
install -Dpm 0755 %{name} %{buildroot}%{_bindir}/%{name}

%check
#

%post
#

%preun
#

%files
%{_bindir}/%{name}


%changelog
* Fri Feb 21 2025 Burak Kirici - 1.4-1
- First release
