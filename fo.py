from os import listdir
from os.path import isfile
from os.path import exists
from os.path import getctime
from os import makedirs
from sys import exit
from sys import argv
from datetime import date
from shutil import copy2
from os import remove

"""
FO or File Organizer is a program meant to take organize

"""
#Global variables for the program
directory = ''
organizationType = 'MONTH'
postCleanup = ''
folderDict = {}

def hasCommandLineArgs():
    """
    Checks if the program was executed with command line argumnets pass in. argv[1] should be the directory and argv[2] should be y/n for the
    postCleanup flag
    """
    if len(argv) == 3:
        if len(argv[1]) != 0:
            if len(argv[2]) !=0:
                print('Command line args found \n' )
                return True
    return False

def validateInput():
    """
    Validates the input whether it was typed by the user or passed in via a command line arg.

    Rules: the directory must exist, and teh postCleanup flag must be 'y' or 'n'
    """
    if not exists(directory):
        print('The directory entered does not exist: ' + directory)
        exit()
    if postCleanup != 'y' and postCleanup != 'n':
        print('You did not enter "y" or "n" you entered: ' + postCleanup)
        exit()

def gatherInput():
    """
    Kicks off the program and gathers the input needed to execute the file organization
    """
    global directory
    global postCleanup
    print('This program will organize files in a given directory by day \n')
    if hasCommandLineArgs():
        directory = argv[1]
        postCleanup = argv[2]
    else:
        directory = input('What directory would you like to organize? ')
        postCleanup = input('Would you like to delete the files afer organization? (y/n): ')
    validateInput()


def copyAndCleanup(fileToCopy):
    copy2(directory + '\\' + fileToCopy, directory + '\\' + key + '\\' + fileToCopy)
    if postCleanup == 'y':
        remove(directory + '\\' + fileToCopy)


#Begin program execution
gatherInput()
filesInDirectory = []
nonFiles = []

for fileName in listdir(directory):
    if isfile(directory + '\\' + fileName):
        filesInDirectory.append(fileName)
    else:
        nonFiles.append(fileName)
if len(nonFiles) != 0:
    print('\nFYI - directories were located in the path provided. The following directories will not be touched as part of this process ' + str(nonFiles))

for singleFile in filesInDirectory:
    timestamp = str(date.fromtimestamp(getctime(directory + '\\' + singleFile)))
    if timestamp in folderDict:
        folderDict[timestamp].append(singleFile);
    else:
        folderDict[timestamp] = [singleFile]

print('\nCreating ' + str(len(folderDict)) + ' new directories' )

for key, fileArray in folderDict.items():
    if exists(directory + '\\' + key):
        for singleFile in fileArray:
            copyAndCleanup(singleFile)
    else:
        makedirs(directory + '\\' + key)
        for singleFile in fileArray:
            copyAndCleanup(singleFile)
