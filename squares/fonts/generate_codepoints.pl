#!/usr/bin/perl -p

# use this script to generate the list of IconData structs from a codepoints list.
# ./generate_codepoints.pl < MaterialIcons-codepoints.txt > ../icon_codepoints.go

BEGIN {
    print("package squares\n")
}

    sub camelize {
        my($name) = @_;

        # $name =~ s/[_-]([a-z])/\U$1/gr;
        # $name =~ s/(?<=[^\A-Z_])_+([^\A-Z_])|([^\A-Z_]+)|_+/\U$1\L$2/g;
        $name =~ s/_(\w)/\U$1/g;
        $name =~ s{^([a-z])}{\U$1}g;
        return $name;
    }

s{([a-z0-9_]+) ([a-f0-9]+)}{"var Icons" . camelize($1) . " = IconData{'\\u$2'}"}e;
