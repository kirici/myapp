Name:           myapp
Version:        1.5.0
Release:        1%{?dist}
Summary:        Placeholder application

License:        GPLv3
Source0:        %{name}-%{version}.tar.gz

BuildRequires:  golang

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
