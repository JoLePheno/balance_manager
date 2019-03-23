CREATE TABLE `Users` (
                      `ID` bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
                      `Name` varchar(200) DEFAULT NULL,
                      `Solde` INT DEFAULT NULL
);
CREATE TABLE `Transactions` (
                        `ID` bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
                        `AccountId` bigint(20) NOT NULL,
                        `Description` varchar(200) DEFAULT NULL,
                        `Currency` varchar(200) DEFAULT NULL,
                        `Notes` varchar(200) DEFAULT NULL,
                        `Amount` int
);