#!/usr/bin/perl
#
#
#
#
#


my $list = qq#a-unix-based-os,11
android,2717
blackberry-os,44
chromeos,27
ios,300
linux,326
fire-os,26
mac,12
mac-os-x,369
windows,3306
windows-mobile,6
windows-phone,40#;


my @devs = split(/\n/,$list);
foreach(@devs){
    chomp;
    my @rows = split(/\,/);
    my $mx   = int($rows[1]);
    my $ua   = $rows[0];
    for (my $i=1; $i <= $mx; $i++){
       print "wget -O- \"https://developers.whatismybrowser.com/useragents/explore/operating_system_name/$ua/$i\" 2>/dev/null | grep -F '/useragents/parse/'|awk -F '>' '{print \$3}'|sed -e '~s#</a##gi' | egrep -v 'Search/Parse' >> $ua-ua-device-list.txt \n";
    }

}
1;
