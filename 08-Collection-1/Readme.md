# Collection 1

1. Extract: [🗂 disk.tar.xz](./01-disk.tar.xz)
2. Run `Autopsy` open `disk.img` and Activate All Analyzes
3. Export Deleted Zip (Xlsx) File: [🗂 f0231232_rels.zip](./Files/f0231232_rels.zip)
4. Change to `.zip` (if it is not!)
5. Extract `unzip f0231232_rels.zip -d f0231232_rels`
6. Search `flag` in Files: `flag = C2 + B6`
```
# xl\SharedStrings.xml

	<si>
		<t>the flag is: {C2 + B6}</t>
	</si>
	<si>
		<t>TUFaQVBBXzA</t>
	</si>
	<si>
		<t>OTE1MWNlYzI</t>
	</si>
	<si>
		<t>MzczYzY4</t>
	</si>

# xl\worksheets\sheet1.xml

t=str
t=s > SharedStrings

			<c r="G8" t="s">
				<v>2</v>
			</c>
```
7. Shared Strings From: `xl\SharedStrings.xml`
```
0: the flag is: {C2 + B6}
1: TUFaQVBBXzA
2: OTE1MWNlYzI
3: MzczYzY4
```

8. Formula and Values for `B6` (Or Create [🗂 Excel File](./08-Collection-1.xlsx))
```
G8=SharedString(2)="OTE1MWNlYzI"
I4="3"
G3=SharedString(3)="MzczYzY4"
H6=CONCATENATE(G8,I4,G3)="OTE1MWNlYzI3MzczYzY4"
E6="0"
C10="0"
B6=CONCATENATE("TUFaQVBBXzA1MmE",E6,"MmI2MzQ5YmVjNTk",C10,H6) ...
... B6="TUFaQVBBXzA1MmE0MmI2MzQ5YmVjNTk0OTE1MWNlYzI3MzczYzY4"
```

9. Decode Base64 String
```
echo "TUFaQVBBXzA1MmE0MmI2MzQ5YmVjNTk0OTE1MWNlYzI3MzczYzY4" | base64 -d
MAZAPA_052a42b6349bec5949151cec27373c68
```